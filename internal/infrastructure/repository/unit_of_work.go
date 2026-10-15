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

package repository

import (
	"context"
	"fmt"

	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/database/postgres/generated"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UnitOfWorkProvider struct {
	pool *pgxpool.Pool
}

func NewUnitOfWorkProvider(pool *pgxpool.Pool) (ports.UnitOfWorkProvider, error) {
	return &UnitOfWorkProvider{pool: pool}, nil
}

func (u *UnitOfWorkProvider) Begin(ctx context.Context) (ports.UnitOfWork, error) {
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	queries := generated.New(tx)
	return NewUnitOfWork(queries, tx), nil
}

func (u *UnitOfWorkProvider) New(ctx context.Context) (ports.UnitOfWork, error) {
	queries := generated.New(u.pool)
	return NewUnitOfWork(queries, nil), nil
}

var _ ports.UnitOfWorkProvider = &UnitOfWorkProvider{}

// UnitOfWork binds every repository to one *generated.Queries. Repositories are
// constructed lazily by their accessors; each is a tiny handle around queries.
type UnitOfWork struct {
	queries   *generated.Queries
	tx        pgx.Tx
	committed bool
}

func NewUnitOfWork(queries *generated.Queries, tx pgx.Tx) ports.UnitOfWork {
	return &UnitOfWork{
		queries: queries,
		tx:      tx,
	}
}

func (u *UnitOfWork) Achievement() repository.AchievementRepository {
	return newPgAchievementRepo(u.queries)
}

func (u *UnitOfWork) AppUser() repository.AppUserRepository {
	return newPgAppUserRepo(u.queries)
}

func (u *UnitOfWork) AppUserCredential() repository.AppUserCredentialRepository {
	return newPgAppUserCredentialRepo(u.queries)
}

func (u *UnitOfWork) AppUserProfile() repository.AppUserProfileRepository {
	return newPgAppUserProfileRepo(u.queries)
}

func (u *UnitOfWork) AppUserSession() repository.AppUserSessionRepository {
	return newPgAppUserSessionRepo(u.queries)
}

func (u *UnitOfWork) AppUserSuspension() repository.AppUserSuspensionRepository {
	return newPgAppUserSuspensionRepo(u.queries)
}

func (u *UnitOfWork) Chat() repository.ChatRepository {
	return newPgChatRepo(u.queries)
}

func (u *UnitOfWork) ChatMember() repository.ChatMemberRepository {
	return newPgChatMemberRepo(u.queries)
}

func (u *UnitOfWork) Course() repository.CourseRepository {
	return newPgCourseRepo(u.queries)
}

func (u *UnitOfWork) CourseLanguage() repository.CourseLanguageRepository {
	return newPgCourseLanguageRepo(u.queries)
}

func (u *UnitOfWork) CourseSubject() repository.CourseSubjectRepository {
	return newPgCourseSubjectRepo(u.queries)
}

func (u *UnitOfWork) CourseSubjectAssessment() repository.CourseSubjectAssessmentRepository {
	return newPgCourseSubjectAssessmentRepo(u.queries)
}

func (u *UnitOfWork) CourseSubjectAssessmentPass() repository.CourseSubjectAssessmentPassRepository {
	return newPgCourseSubjectAssessmentPassRepo(u.queries)
}

func (u *UnitOfWork) CourseSubjectAssessmentQuestion() repository.CourseSubjectAssessmentQuestionRepository {
	return newPgCourseSubjectAssessmentQuestionRepo(u.queries)
}

func (u *UnitOfWork) CourseSubjectDoc() repository.CourseSubjectDocRepository {
	return newPgCourseSubjectDocRepo(u.queries)
}

func (u *UnitOfWork) CourseSubjectDocMd() repository.CourseSubjectDocMdRepository {
	return newPgCourseSubjectDocMdRepo(u.queries)
}

func (u *UnitOfWork) CourseSubjectItem() repository.CourseSubjectItemRepository {
	return newPgCourseSubjectItemRepo(u.queries)
}

func (u *UnitOfWork) DailyActivity() repository.DailyActivityRepository {
	return newPgDailyActivityRepo(u.queries)
}

func (u *UnitOfWork) Inventory() repository.InventoryRepository {
	return newPgInventoryRepo(u.queries)
}

func (u *UnitOfWork) InventorySlot() repository.InventorySlotRepository {
	return newPgInventorySlotRepo(u.queries)
}

func (u *UnitOfWork) Item() repository.ItemRepository {
	return newPgItemRepo(u.queries)
}

func (u *UnitOfWork) Language() repository.LanguageRepository {
	return newPgLanguageRepo(u.queries)
}

func (u *UnitOfWork) Message() repository.MessageRepository {
	return newPgMessageRepo(u.queries)
}

func (u *UnitOfWork) Notification() repository.NotificationRepository {
	return newPgNotificationRepo(u.queries)
}

func (u *UnitOfWork) RegistrationInvite() repository.RegistrationInviteRepository {
	return newPgRegistrationInviteRepo(u.queries)
}

func (u *UnitOfWork) StudentAchievement() repository.StudentAchievementRepository {
	return newPgStudentAchievementRepo(u.queries)
}

func (u *UnitOfWork) StudentCourseStats() repository.StudentCourseStatsRepository {
	return newPgStudentCourseStatsRepo(u.queries)
}

func (u *UnitOfWork) StudentCourseSubject() repository.StudentCourseSubjectRepository {
	return newPgStudentCourseSubjectRepo(u.queries)
}

func (u *UnitOfWork) StudentCourseSubjectAssessmentAttempt() repository.StudentCourseSubjectAssessmentAttemptRepository {
	return newPgStudentCourseSubjectAssessmentAttemptRepo(u.queries)
}

func (u *UnitOfWork) StudentCourseSubjectSubmissionAttempt() repository.StudentCourseSubjectSubmissionAttemptRepository {
	return newPgStudentCourseSubjectSubmissionAttemptRepo(u.queries)
}

func (u *UnitOfWork) StudentGlobalStats() repository.StudentGlobalStatsRepository {
	return newPgStudentGlobalStatsRepo(u.queries)
}

func (u *UnitOfWork) StudentGroup() repository.StudentGroupRepository {
	return newPgStudentGroupRepo(u.queries)
}

func (u *UnitOfWork) StudentGroupClass() repository.StudentGroupClassRepository {
	return newPgStudentGroupClassRepo(u.queries)
}

func (u *UnitOfWork) StudentGroupCourse() repository.StudentGroupCourseRepository {
	return newPgStudentGroupCourseRepo(u.queries)
}

func (u *UnitOfWork) StudentGroupMember() repository.StudentGroupMemberRepository {
	return newPgStudentGroupMemberRepo(u.queries)
}

func (u *UnitOfWork) StudentGroupMemberClassAttendance() repository.StudentGroupMemberClassAttendanceRepository {
	return newPgStudentGroupMemberClassAttendanceRepo(u.queries)
}

func (u *UnitOfWork) TOTPSecret() repository.TOTPSecretRepository {
	return newPgTOTPSecretRepo(u.queries)
}

func (u *UnitOfWork) Teacher() repository.TeacherRepository {
	return newPgTeacherRepo(u.queries)
}

func (u *UnitOfWork) TeacherGroup() repository.TeacherGroupRepository {
	return newPgTeacherGroupRepo(u.queries)
}

func (u *UnitOfWork) Commit(ctx context.Context) error {
	if u.tx == nil {
		return nil
	}
	if u.committed {
		return fmt.Errorf("transaction already committed")
	}
	err := u.tx.Commit(ctx)
	if err == nil {
		u.committed = true
	}
	return err
}

func (u *UnitOfWork) Rollback(ctx context.Context) error {
	if u.tx == nil {
		return nil
	}
	if u.committed {
		return nil
	}
	return u.tx.Rollback(ctx)
}

var _ ports.UnitOfWork = &UnitOfWork{}
