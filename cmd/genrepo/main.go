// Command genrepo generates the boilerplate CRUD repository adapters for the
// Brainiac persistence layer.
//
// It exists to remove the ~400 lines of hand-written glue that each entity
// previously required. For every entity listed below it emits a
// zz_generated_<file>.go adapter that wires the sqlc-generated *Queries methods
// (Get<Name>, Create<Name>, Save<Name>, Update<Name>, Delete<Name>) to the
// domain entity through a compile-time struct conversion.
//
// The output is plain, typed Go — no runtime reflection. Bespoke queries
// (GetByUsername, ExistsByEmail, ...) remain hand-written next to the generated
// file in the same package.
//
// Model is only set when sqlc's struct name differs from the domain entity name
// (sqlc singularises table names, so e.g. table student_course_stats yields
// generated.StudentCourseStat while the entity is entity.StudentCourseStats).
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
)

type entitySpec struct {
	Name     string   // "AppUser" — used as entity.<Name> and as the query prefix
	File     string   // base filename without the zz_generated_ prefix and .go suffix
	Model    string   // generated struct name if it differs from Name (else empty)
	Writable []string // fields copied into Create/Save/Update params (column == field name)
}

var specs = []entitySpec{
	{
		Name:     "Achievement",
		File:     "achievement",
		Writable: []string{"GUID", "TitleI18n", "DescriptionI18n", "ResourceUrl", "IsHidden", "Meta", "Conditions", "RewardExperience", "RewardMoney", "PublishedAt"},
	},
	{
		Name:     "AppUser",
		File:     "appuser",
		Writable: []string{"GUID", "Username", "ActivatedAt"},
	},
	{
		Name:     "AppUserCredential",
		File:     "appuser_credential",
		Writable: []string{"AppUserGUID", "Email", "PasswordHash"},
	},
	{
		Name:     "AppUserProfile",
		File:     "appuser_profile",
		Writable: []string{"AppUserGUID", "Name", "Surname", "Patronymic", "Nickname", "Bio", "PreferredLanguage", "ProfileDiscovery", "AvatarUrl", "EditingLockedAt"},
	},
	{
		Name:     "AppUserSession",
		File:     "appuser_session",
		Writable: []string{"GUID", "AppUserGUID", "LastIPV4", "LastIPV6", "LastAgent", "LastSeenAt", "ExpireAt"},
	},
	{
		Name:     "AppUserSuspension",
		File:     "app_user_suspension",
		Writable: []string{"GUID", "AppUserGUID", "Reason", "ExpireAt", "ExpiredAt"},
	},
	{
		Name:     "Chat",
		File:     "chat",
		Writable: []string{"GUID", "Title", "Topic", "AvatarUrl", "StudentsCanText", "StudentsSeeMembers", "PreserveMessages", "Meta"},
	},
	{
		Name:     "ChatMember",
		File:     "chat_member",
		Writable: []string{"GUID", "AppUserGUID", "ChatGUID", "ChatRole", "Name", "CustomRoleName", "LeftAt", "SeenAt"},
	},
	{
		Name:     "Course",
		File:     "course",
		Writable: []string{"GUID", "TitleI18n", "DescriptionI18n", "Style", "Meta", "PublishedAt"},
	},
	{
		Name:     "CourseLanguage",
		File:     "course_language",
		Writable: []string{"GUID", "CourseGUID", "LanguageGUID", "CalculatedSupportPercentage", "IsBeta", "IsNew"},
	},
	{
		Name:     "CourseSubject",
		File:     "course_subject",
		Writable: []string{"GUID", "CourseGUID", "ParentGUID", "TitleI18n", "DescriptionI18n", "RewardExperience", "CompletionCondition", "HideChildren", "Style", "Meta", "PublishedAt"},
	},
	{
		Name:     "CourseSubjectAssessment",
		File:     "course_subject_assessment",
		Writable: []string{"GUID", "CourseSubjectGUID", "TitleI18n", "DescriptionI18n", "IsAttestation", "HideAnswers", "HideAnswerValidity", "HideScore", "UseTimedQuestions", "MaxTime", "MaxAttempts", "ScoringMethod", "RequiredScore", "ShortQuestionsCount", "NormalQuestionsCount", "LongQuestionsCount", "PublishedAt"},
	},
	{
		Name:     "CourseSubjectAssessmentPass",
		File:     "course_subject_assessment_pass",
		Writable: []string{"GUID", "CourseSubjectAssessmentGUID", "AppUserGUID", "IsActive", "AttemptsLeft", "ExpireAt"},
	},
	{
		Name:     "CourseSubjectAssessmentQuestion",
		File:     "course_subject_assessment_question",
		Writable: []string{"GUID", "CourseSubjectAssessmentGUID", "TitleI18n", "DescriptionI18n", "QuestionType", "AttachmentUrl", "Options", "CorrectOptions", "MaxCorrectOptions", "IsMultipleChoice", "MaxOptions", "MaxAnswerTime", "UseTextAnswer", "CorrectTextAnswer", "ExampleUrl", "PublishedAt"},
	},
	{
		Name:     "CourseSubjectDoc",
		File:     "course_subject_doc",
		Writable: []string{"GUID", "CourseSubjectGUID", "TitleI18n", "DescriptionI18n", "SubjectDocType", "ExampleMeta", "SubmissionMeta", "LiteratureMeta", "DocIndex", "PublishedAt"},
	},
	{
		Name:     "CourseSubjectDocMd",
		File:     "course_subject_doc_md",
		Writable: []string{"GUID", "CourseSubjectDocGUID", "LanguageCode", "Md", "PublishedAt"},
	},
	{
		Name:     "CourseSubjectItem",
		File:     "course_subject_item",
		Writable: []string{"GUID", "CourseSubjectGUID", "ItemGUID", "Amount"},
	},
	{
		Name:     "DailyActivity",
		File:     "daily_activity",
		Writable: []string{"GUID", "AppUserGUID", "Day", "ExperienceEarned", "LevelsEarned", "SubjectsCompleted", "AssessmentsCompleted", "PracticesCompleted"},
	},
	{
		Name:     "Inventory",
		File:     "inventory",
		Writable: []string{"AppUserGUID", "MaxSlots"},
	},
	{
		Name:     "InventorySlot",
		File:     "inventory_slot",
		Writable: []string{"GUID", "InventoryGUID", "ItemGUID", "Amount", "AmountInExchange", "ExpireAt"},
	},
	{
		Name:     "Item",
		File:     "item",
		Writable: []string{"GUID", "TitleI18n", "DescriptionI18n", "ResourceUrl", "Meta", "StackSize", "Rarity", "AllowExchange", "ShopPrice", "OneTimePurchase", "ShopQuantity", "LevelRequired", "OnSaleSince", "OnSaleUntil", "PublishedAt"},
	},
	{
		Name:     "Language",
		File:     "language",
		Writable: []string{"GUID", "LanguageCode", "EnglishTitle", "LocalTitle", "Emoji", "IsSupported", "IsBeta"},
	},
	{
		Name:     "Message",
		File:     "message",
		Writable: []string{"GUID", "ChatMemberGUID", "ChatGUID", "ResourceUrl", "Content", "Urgency", "Meta", "SeenBy", "EditedAt", "PinnedAt", "ExpireAt"},
	},
	{
		Name:     "Notification",
		File:     "notification",
		Writable: []string{"GUID", "AppUserGUID", "TitleI18n", "ContentI18n", "ResourceUrl", "Urgency", "Meta", "SeenAt", "ExpireAt"},
	},
	{
		Name:     "RegistrationInvite",
		File:     "registration_invite",
		Writable: []string{"AppUserGUID", "InvitedByAppUserGUID", "InviteCode", "Email", "Message", "Name", "Surname", "Patronymic", "Nickname", "Username", "StudentGroups", "TeacherGroups", "Chats", "ExpireAt", "UsedAt"},
	},
	{
		Name:     "StudentAchievement",
		File:     "student_achievement",
		Writable: []string{"GUID", "AppUserGUID", "AchievementGUID"},
	},
	{
		Name:     "StudentCourseStats",
		File:     "student_course_stats",
		Model:    "StudentCourseStat",
		Writable: []string{"GUID", "AppUserGUID", "CourseGUID", "Experience", "Level", "Meta"},
	},
	{
		Name:     "StudentCourseSubject",
		File:     "student_course_subject",
		Writable: []string{"GUID", "AppUserGUID", "CourseSubjectGUID", "CompletedAt", "SubjectStatus", "IsFavorite", "Meta"},
	},
	{
		Name:     "StudentCourseSubjectAssessmentAttempt",
		File:     "student_course_subject_assessment_attempt",
		Writable: []string{"GUID", "AppUserGUID", "CourseSubjectAssessmentGUID", "AttemptStatus", "Score", "IsSuccess", "Answers"},
	},
	{
		Name:     "StudentCourseSubjectSubmissionAttempt",
		File:     "student_course_subject_submission_attempt",
		Writable: []string{"GUID", "AppUserGUID", "CourseSubjectDocGUID", "SubmissionStatus", "Files", "Meta", "StudentComment", "TeacherComment"},
	},
	{
		Name:     "StudentGlobalStats",
		File:     "student_global_stats",
		Model:    "StudentGlobalStat",
		Writable: []string{"AppUserGUID", "Level", "OtherExperience", "Money", "MoneyEarned", "SubjectsCompleted", "AssessmentsCompleted", "AssessmentsFailed", "AssessmentsTerminated", "WorksSubmitted", "ItemsCollected", "ItemsUsed", "ItemsSold", "ItemsExchanged", "BoxesOpened", "MaxDailyStreak", "MessagesSent"},
	},
	{
		Name:     "StudentGroup",
		File:     "student_group",
		Writable: []string{"GUID", "Title", "Description", "EnableGamification"},
	},
	{
		Name:     "StudentGroupClass",
		File:     "student_group_class",
		Writable: []string{"GUID", "StudentGroupGUID", "Title", "StartTime", "EndTime", "Objectives", "AttestationAssessmentsEnabled", "AttendanceClosedAt", "MarksApprovedAt", "TOTPSecret", "Room", "ClassType"},
	},
	{
		Name:     "StudentGroupCourse",
		File:     "student_group_course",
		Writable: []string{"GUID", "StudentGroupGUID", "CourseGUID", "FullAccess", "CloseSubmissions"},
	},
	{
		Name:     "StudentGroupMember",
		File:     "student_group_member",
		Writable: []string{"GUID", "AppUserGUID", "StudentGroupGUID", "StudentRole", "Notes", "Team", "TeamResponsibilities"},
	},
	{
		Name:     "StudentGroupMemberClassAttendance",
		File:     "student_group_member_class_attendance",
		Writable: []string{"GUID", "StudentGroupMemberGUID", "StudentGroupClassGUID", "FirstSeenAt", "LastSeenAt", "PresentAt", "ObjectivesCompletion", "Comment", "Grade"},
	},
	{
		Name:     "TOTPSecret",
		File:     "totp_secret",
		Writable: []string{"AppUserGUID", "SecretBase32", "LastUsedAt"},
	},
	{
		Name:     "Teacher",
		File:     "teacher",
		Writable: []string{"AppUserGUID", "ActiveSince", "ManageStudents", "ManageAttendance", "ManageResults", "ManageCourses", "IsAdmin", "ContactInfo"},
	},
	{
		Name:     "TeacherGroup",
		File:     "teacher_group",
		Writable: []string{"GUID", "TeacherGUID", "StudentGroupGUID", "ActiveSince", "ManageAttendance", "ManageResults", "ManageStudents", "ManageAchievements", "TeacherUntil"},
	},
}

func main() {
	out := flag.String("out", ".", "directory to write generated files into")
	flag.Parse()

	for _, s := range specs {
		src, err := render(s)
		if err != nil {
			log.Fatalf("render %s: %v", s.Name, err)
		}
		path := filepath.Join(*out, "zz_generated_"+s.File+".go")
		if err := os.WriteFile(path, src, 0o644); err != nil {
			log.Fatalf("write %s: %v", path, err)
		}
		fmt.Fprintf(os.Stderr, "generated %s\n", path)
	}
}

func render(s entitySpec) ([]byte, error) {
	model := s.Model
	if model == "" {
		model = s.Name
	}

	var b bytes.Buffer

	structName := "pg" + s.Name + "Repo"
	ctorName := "newPg" + s.Name + "Repo"

	b.WriteString("// Code generated by genrepo. DO NOT EDIT.\n\n")
	b.WriteString("package repository\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"context\"\n\n")
	b.WriteString("\t\"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/database/postgres/generated\"\n")
	b.WriteString("\t\"github.com/ReallyWeirdCat/brainiac/pkg/domain/entity\"\n")
	b.WriteString("\trepo \"github.com/ReallyWeirdCat/brainiac/pkg/domain/repository\"\n")
	b.WriteString("\t\"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject\"\n")
	b.WriteString(")\n\n")

	fmt.Fprintf(&b, "type %s struct {\n\tq *generated.Queries\n}\n\n", structName)
	fmt.Fprintf(&b, "func %s(q *generated.Queries) repo.%sRepository {\n\treturn &%s{q: q}\n}\n\n", ctorName, s.Name, structName)

	// Get
	fmt.Fprintf(&b, "func (r *%s) Get(ctx context.Context, guid valueobject.GUID) (*entity.%s, error) {\n", structName, s.Name)
	fmt.Fprintf(&b, "\tm, err := r.q.Get%s(ctx, guid)\n", s.Name)
	b.WriteString("\tif err != nil {\n\t\treturn nil, MapNotFound(err)\n\t}\n")
	fmt.Fprintf(&b, "\treturn ptr(entity.%s(m)), nil\n}\n\n", s.Name)

	// Create / Save / Update share the param-mapping body.
	for _, op := range []string{"Create", "Save", "Update"} {
		fmt.Fprintf(&b, "func (r *%s) %s(ctx context.Context, obj *entity.%s) (*entity.%s, error) {\n", structName, op, s.Name, s.Name)
		fmt.Fprintf(&b, "\tm, err := r.q.%s%s(ctx, generated.%s%sParams{\n", op, s.Name, op, s.Name)
		for _, f := range s.Writable {
			fmt.Fprintf(&b, "\t\t%s: obj.%s,\n", f, f)
		}
		b.WriteString("\t})\n")
		b.WriteString("\tif err != nil {\n")
		if op == "Update" {
			b.WriteString("\t\treturn nil, MapNotFound(err)\n")
		} else {
			b.WriteString("\t\treturn nil, err\n")
		}
		b.WriteString("\t}\n")
		fmt.Fprintf(&b, "\treturn ptr(entity.%s(m)), nil\n}\n\n", s.Name)
	}

	// Delete
	fmt.Fprintf(&b, "func (r *%s) Delete(ctx context.Context, guid valueobject.GUID) error {\n", structName)
	fmt.Fprintf(&b, "\treturn MapNotFound(r.q.Delete%s(ctx, guid))\n}\n", s.Name)

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		// Return the unformatted source so the user can inspect what failed.
		return nil, fmt.Errorf("format source: %w\n%s", err, b.String())
	}
	return formatted, nil
}
