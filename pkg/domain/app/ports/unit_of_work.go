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

package ports

import (
	"context"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
)

type UnitOfWorkProvider interface {
	// Returns an instance of UnitOfWork with a transaction
	Begin(ctx context.Context) (UnitOfWork, error)
	// Returns an instance of UnitOfWork with no transaction
	New(ctx context.Context) (UnitOfWork, error)
}

// UnitOfWork exposes every repository for a single transaction (or a
// non-transactional snapshot of the pool).
type UnitOfWork interface {
	Achievement() repository.AchievementRepository
	AppUser() repository.AppUserRepository
	AppUserCredential() repository.AppUserCredentialRepository
	AppUserProfile() repository.AppUserProfileRepository
	AppUserSession() repository.AppUserSessionRepository
	AppUserSuspension() repository.AppUserSuspensionRepository
	Chat() repository.ChatRepository
	ChatMember() repository.ChatMemberRepository
	Course() repository.CourseRepository
	CourseLanguage() repository.CourseLanguageRepository
	CourseSubject() repository.CourseSubjectRepository
	CourseSubjectAssessment() repository.CourseSubjectAssessmentRepository
	CourseSubjectAssessmentPass() repository.CourseSubjectAssessmentPassRepository
	CourseSubjectAssessmentQuestion() repository.CourseSubjectAssessmentQuestionRepository
	CourseSubjectDoc() repository.CourseSubjectDocRepository
	CourseSubjectDocMd() repository.CourseSubjectDocMdRepository
	CourseSubjectItem() repository.CourseSubjectItemRepository
	DailyActivity() repository.DailyActivityRepository
	Inventory() repository.InventoryRepository
	InventorySlot() repository.InventorySlotRepository
	Item() repository.ItemRepository
	Language() repository.LanguageRepository
	Message() repository.MessageRepository
	Notification() repository.NotificationRepository
	RegistrationInvite() repository.RegistrationInviteRepository
	StudentAchievement() repository.StudentAchievementRepository
	StudentCourseStats() repository.StudentCourseStatsRepository
	StudentCourseSubject() repository.StudentCourseSubjectRepository
	StudentCourseSubjectAssessmentAttempt() repository.StudentCourseSubjectAssessmentAttemptRepository
	StudentCourseSubjectSubmissionAttempt() repository.StudentCourseSubjectSubmissionAttemptRepository
	StudentGlobalStats() repository.StudentGlobalStatsRepository
	StudentGroup() repository.StudentGroupRepository
	StudentGroupClass() repository.StudentGroupClassRepository
	StudentGroupCourse() repository.StudentGroupCourseRepository
	StudentGroupMember() repository.StudentGroupMemberRepository
	StudentGroupMemberClassAttendance() repository.StudentGroupMemberClassAttendanceRepository
	TOTPSecret() repository.TOTPSecretRepository
	Teacher() repository.TeacherRepository
	TeacherGroup() repository.TeacherGroupRepository
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
