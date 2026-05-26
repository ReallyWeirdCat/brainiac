// Copyright (c) 2026 Nikolai Papin
//
// This file is part of Brainiac gamification and education platform
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package shared

import (
	"fmt"
)

// QuestionTypeEnum represents the type of question (0=short_question, 1=normal_question, 2=long_question)
type QuestionTypeEnum int16

const (
	ShortQuestion  QuestionTypeEnum = 0
	NormalQuestion QuestionTypeEnum = 1
	LongQuestion   QuestionTypeEnum = 2
)

func (e QuestionTypeEnum) String() string {
	switch e {
	case ShortQuestion:
		return "short_question"
	case NormalQuestion:
		return "normal_question"
	case LongQuestion:
		return "long_question"
	default:
		return "unknown"
	}
}

func (e QuestionTypeEnum) Valid() bool {
	return e >= 0 && e <= 2
}

// ScoringMethodEnum represents the scoring method (0=max_score, 1=average_score, 2=latest_score)
type ScoringMethodEnum int16

const (
	MaxScore     ScoringMethodEnum = 0
	AverageScore ScoringMethodEnum = 1
	LatestScore  ScoringMethodEnum = 2
)

func (e ScoringMethodEnum) String() string {
	switch e {
	case MaxScore:
		return "max_score"
	case AverageScore:
		return "average_score"
	case LatestScore:
		return "latest_score"
	default:
		return "unknown"
	}
}

func (e ScoringMethodEnum) Valid() bool {
	return e >= 0 && e <= 2
}

// CompletionConditionEnum represents completion condition (0=manual_completion, 1=attestation_test, 2=approved_practice, 3=all)
type CompletionConditionEnum int16

const (
	ManualCompletion    CompletionConditionEnum = 0
	AttestationTest     CompletionConditionEnum = 1
	ApprovedPractice    CompletionConditionEnum = 2
	AllCompletion       CompletionConditionEnum = 3
)

func (e CompletionConditionEnum) String() string {
	switch e {
	case ManualCompletion:
		return "manual_completion"
	case AttestationTest:
		return "attestation_test"
	case ApprovedPractice:
		return "approved_practice"
	case AllCompletion:
		return "all"
	default:
		return "unknown"
	}
}

func (e CompletionConditionEnum) Valid() bool {
	return e >= 0 && e <= 3
}

// SubjectStatusEnum represents subject status (0=untouched, 1=in_progress, 2=forgotten, 3=done)
type SubjectStatusEnum int16

const (
	Untouched  SubjectStatusEnum = 0
	InProgress SubjectStatusEnum = 1
	Forgotten  SubjectStatusEnum = 2
	Done       SubjectStatusEnum = 3
)

func (e SubjectStatusEnum) String() string {
	switch e {
	case Untouched:
		return "untouched"
	case InProgress:
		return "in_progress"
	case Forgotten:
		return "forgotten"
	case Done:
		return "done"
	default:
		return "unknown"
	}
}

func (e SubjectStatusEnum) Valid() bool {
	return e >= 0 && e <= 3
}

// GradeEnum represents grade (0=unspecified, 1=very_poor, 2=poor, 3=satisfactory, 4=good, 5=very_good)
type GradeEnum int16

const (
	Unspecified  GradeEnum = 0
	VeryPoor     GradeEnum = 1
	Poor         GradeEnum = 2
	Satisfactory GradeEnum = 3
	Good         GradeEnum = 4
	VeryGood     GradeEnum = 5
)

func (e GradeEnum) String() string {
	switch e {
	case Unspecified:
		return "unspecified"
	case VeryPoor:
		return "very_poor"
	case Poor:
		return "poor"
	case Satisfactory:
		return "satisfactory"
	case Good:
		return "good"
	case VeryGood:
		return "very_good"
	default:
		return "unknown"
	}
}

func (e GradeEnum) Valid() bool {
	return e >= 0 && e <= 5
}

// SubjectDocTypeEnum represents subject document type (0=theory, 1=practice, 2=example, 3=literature)
type SubjectDocTypeEnum int16

const (
	TheoryDoc     SubjectDocTypeEnum = 0
	PracticeDoc   SubjectDocTypeEnum = 1
	ExampleDoc    SubjectDocTypeEnum = 2
	LiteratureDoc SubjectDocTypeEnum = 3
)

func (e SubjectDocTypeEnum) String() string {
	switch e {
	case TheoryDoc:
		return "theory"
	case PracticeDoc:
		return "practice"
	case ExampleDoc:
		return "example"
	case LiteratureDoc:
		return "literature"
	default:
		return "unknown"
	}
}

func (e SubjectDocTypeEnum) Valid() bool {
	return e >= 0 && e <= 3
}

// AttemptStatusEnum represents attempt status (0=in_progress, 1=finished, 2=flagged)
type AttemptStatusEnum int16

const (
	AttemptInProgress AttemptStatusEnum = 0
	AttemptFinished   AttemptStatusEnum = 1
	AttemptFlagged    AttemptStatusEnum = 2
)

func (e AttemptStatusEnum) String() string {
	switch e {
	case AttemptInProgress:
		return "in_progress"
	case AttemptFinished:
		return "finished"
	case AttemptFlagged:
		return "flagged"
	default:
		return "unknown"
	}
}

func (e AttemptStatusEnum) Valid() bool {
	return e >= 0 && e <= 2
}

// ClassTypeEnum represents class type (0=practice, 1=lecture, 2=attestation, 3=consult)
type ClassTypeEnum int16

const (
	PracticeClass   ClassTypeEnum = 0
	LectureClass    ClassTypeEnum = 1
	AttestationClass ClassTypeEnum = 2
	ConsultClass    ClassTypeEnum = 3
)

func (e ClassTypeEnum) String() string {
	switch e {
	case PracticeClass:
		return "practice"
	case LectureClass:
		return "lecture"
	case AttestationClass:
		return "attestation"
	case ConsultClass:
		return "consult"
	default:
		return "unknown"
	}
}

func (e ClassTypeEnum) Valid() bool {
	return e >= 0 && e <= 3
}

// ChatRoleEnum represents chat role (0=member, 1=admin, 2=owner)
type ChatRoleEnum int16

const (
	Member ChatRoleEnum = 0
	Admin  ChatRoleEnum = 1
	Owner  ChatRoleEnum = 2
)

func (e ChatRoleEnum) String() string {
	switch e {
	case Member:
		return "member"
	case Admin:
		return "admin"
	case Owner:
		return "owner"
	default:
		return "unknown"
	}
}

func (e ChatRoleEnum) Valid() bool {
	return e >= 0 && e <= 2
}

// UrgencyEnum represents urgency level (0=low, 1=normal, 2=important)
type UrgencyEnum int16

const (
	LowUrgency      UrgencyEnum = 0
	NormalUrgency   UrgencyEnum = 1
	ImportantUrgency UrgencyEnum = 2
)

func (e UrgencyEnum) String() string {
	switch e {
	case LowUrgency:
		return "low"
	case NormalUrgency:
		return "normal"
	case ImportantUrgency:
		return "important"
	default:
		return "unknown"
	}
}

func (e UrgencyEnum) Valid() bool {
	return e >= 0 && e <= 2
}

// RarityEnum represents rarity level (0=common, 1=uncommon, 2=rare, 3=epic, 4=legendary, 5=mythical)
type RarityEnum int16

const (
	CommonRarity    RarityEnum = 0
	UncommonRarity  RarityEnum = 1
	RareRarity      RarityEnum = 2
	EpicRarity      RarityEnum = 3
	LegendaryRarity RarityEnum = 4
	MythicalRarity  RarityEnum = 5
)

func (e RarityEnum) String() string {
	switch e {
	case CommonRarity:
		return "common"
	case UncommonRarity:
		return "uncommon"
	case RareRarity:
		return "rare"
	case EpicRarity:
		return "epic"
	case LegendaryRarity:
		return "legendary"
	case MythicalRarity:
		return "mythical"
	default:
		return "unknown"
	}
}

func (e RarityEnum) Valid() bool {
	return e >= 0 && e <= 5
}

// ProfileDiscoveryEnum represents profile discovery level (0=admins_only, 1=teachers, 2=group_mates, 3=course_mates, 4=users, 5=guests)
type ProfileDiscoveryEnum int16

const (
	AdminsOnly   ProfileDiscoveryEnum = 0
	Teachers     ProfileDiscoveryEnum = 1
	GroupMates   ProfileDiscoveryEnum = 2
	CourseMates  ProfileDiscoveryEnum = 3
	AllUsers     ProfileDiscoveryEnum = 4
	Guests       ProfileDiscoveryEnum = 5
)

func (e ProfileDiscoveryEnum) String() string {
	switch e {
	case AdminsOnly:
		return "admins_only"
	case Teachers:
		return "teachers"
	case GroupMates:
		return "group_mates"
	case CourseMates:
		return "course_mates"
	case AllUsers:
		return "users"
	case Guests:
		return "guests"
	default:
		return "unknown"
	}
}

func (e ProfileDiscoveryEnum) Valid() bool {
	return e >= 0 && e <= 5
}

// StudentRoleEnum represents student role (0=unspecified, 1=student, 2=basic, 3=intermediate)
type StudentRoleEnum int16

const (
	UnspecifiedRole StudentRoleEnum = 0
	StudentRole     StudentRoleEnum = 1
	BasicRole       StudentRoleEnum = 2
	IntermediateRole StudentRoleEnum = 3
)

func (e StudentRoleEnum) String() string {
	switch e {
	case UnspecifiedRole:
		return "unspecified"
	case StudentRole:
		return "student"
	case BasicRole:
		return "basic"
	case IntermediateRole:
		return "intermediate"
	default:
		return "unknown"
	}
}

func (e StudentRoleEnum) Valid() bool {
	return e >= 0 && e <= 3
}

// ParseQuestionTypeEnum parses a string into a QuestionTypeEnum
func ParseQuestionTypeEnum(s string) (QuestionTypeEnum, error) {
	switch s {
	case "short_question":
		return ShortQuestion, nil
	case "normal_question":
		return NormalQuestion, nil
	case "long_question":
		return LongQuestion, nil
	default:
		return -1, fmt.Errorf("invalid question type: %s", s)
	}
}

// ParseScoringMethodEnum parses a string into a ScoringMethodEnum
func ParseScoringMethodEnum(s string) (ScoringMethodEnum, error) {
	switch s {
	case "max_score":
		return MaxScore, nil
	case "average_score":
		return AverageScore, nil
	case "latest_score":
		return LatestScore, nil
	default:
		return -1, fmt.Errorf("invalid scoring method: %s", s)
	}
}

// ParseGradeEnum parses a string into a GradeEnum
func ParseGradeEnum(s string) (GradeEnum, error) {
	switch s {
	case "unspecified":
		return Unspecified, nil
	case "very_poor":
		return VeryPoor, nil
	case "poor":
		return Poor, nil
	case "satisfactory":
		return Satisfactory, nil
	case "good":
		return Good, nil
	case "very_good":
		return VeryGood, nil
	default:
		return -1, fmt.Errorf("invalid grade: %s", s)
	}
}

// Value returns the integer value for database operations
func (e QuestionTypeEnum) Value() int16 {
	return int16(e)
}

// Value returns the integer value for database operations
func (e ScoringMethodEnum) Value() int16 {
	return int16(e)
}

// Value returns the integer value for database operations
func (e CompletionConditionEnum) Value() int16 {
	return int16(e)
}

// Value returns the integer value for database operations
func (e SubjectStatusEnum) Value() int16 {
	return int16(e)
}

// Value returns the integer value for database operations
func (e GradeEnum) Value() int16 {
	return int16(e)
}

// Value returns the integer value for database operations
func (e SubjectDocTypeEnum) Value() int16 {
	return int16(e)
}

// Value returns the integer value for database operations
func (e AttemptStatusEnum) Value() int16 {
	return int16(e)
}

// Value returns the integer value for database operations
func (e ClassTypeEnum) Value() int16 {
	return int16(e)
}

// Value returns the integer value for database operations
func (e ChatRoleEnum) Value() int16 {
	return int16(e)
}

// Value returns the integer value for database operations
func (e UrgencyEnum) Value() int16 {
	return int16(e)
}

// Value returns the integer value for database operations
func (e RarityEnum) Value() int16 {
	return int16(e)
}

// Value returns the integer value for database operations
func (e ProfileDiscoveryEnum) Value() int16 {
	return int16(e)
}

// Value returns the integer value for database operations
func (e StudentRoleEnum) Value() int16 {
	return int16(e)
}
