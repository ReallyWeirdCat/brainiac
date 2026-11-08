# Brainiac domain

Business rules for the Brainiac education platform. For engineering conventions, layout and commands, see [AGENTS.md](./AGENTS.md).

## Glossary

- **Subject (module)**: a unit of a course covering a specific topic.
- **Attestation assessment**: an assessment that can only be attempted under supervision or with permission.
- **Assessment pass**: a teacher's permission for selected students to start an attestation assessment.
- **Self-attendance**: a student marks themselves present by scanning a QR code during a class.
- **Money points**: virtual currency used to buy items; some items can be converted back into it.

## Courses

Teachers create course subjects (modules) that feature specific topics. Subjects are arranged in a graph and can branch. This can be a single branch in a linear course, or a tree with many branches where students choose what to learn. Some subjects are locked until their parent subjects are finished by the student.

Subjects can be completed in several ways:
- The student marks completion manually (when the module does not test the student).
- The student completes the task or assessment attached to the subject.

Some subjects have an assessment marked as attestation, meaning the student can only start an attempt with the teacher's permission (under supposed supervision).

## Course languages

Course languages are different from interface languages. Languages can be created and implemented per course, so one course can serve people of different languages. This also applies to assessments, questions and tasks. A translation percentage can be calculated to display language support for a course.

## Classes

Classes are scheduled for a certain time and feature attendance logging. An attendance record has these fields:

- `PresentAt`: latest mark, by self-attendance or by the teacher.
- `FirstSeenAt`: first time the student was marked present during the class.
- `LastSeenAt`: last time the student was marked present during the class.

The teacher marks students manually or via a self-attendance QR code. `FirstSeenAt` and `LastSeenAt` allow marking attendance twice (at the start and at the end) to show how long the person was present and whether they left. `PresentAt` acts like `LastSeenAt` but can be edited by the teacher.

**`PresentAt` is the only field that matters to the system.** If it is null, the student is considered absent, even if `FirstSeenAt` and `LastSeenAt` are set.

A TOTP secret on the class generates the self-attendance QR codes, each active for a brief moment.

Class types (`practice`, `lecture`, `attestation`, `consult`) are cosmetic and do not affect system behavior.

## Tasks

Task submissions are checked and rated by the teacher (possibly with AI assistance in the future). They usually require submitting code, a report, text or other information.

## Assessments

An assessment may or may not be an attestation assessment.

**Non-attestation:** completion is not mandatory and needs no supervision. These let students test their own knowledge.

**Attestation:** an attempt can only be started under certain conditions:
- The teacher opens the test for selected students (assigns an assessment pass), or
- Attestation assessments are enabled for an active class and the student is marked present (`PresentAt` set). The student can then start an attempt without the teacher's permission.

Attestation assessments include anti-cheating mechanisms, especially for when the teacher is not present:
- Timed questions leave no time to search or ask AI assistants.
- Questions may be rendered on the backend as noisy images or SVGs, so they cannot be copy-pasted or read by AI vision.

## Messenger

The messenger has chat rooms for teachers and students. Rooms may be restricted to hide other students from the student and disallow student messages, which turns the group into a channel.

Students who can see each other can message each other; a room is created for the two of them. In a direct-message chat, both are admins of the room.

## Teachers

Teacher permissions:

| Permission | Meaning |
|---|---|
| `ManageStudents` | Change profiles of students in their groups |
| `ManageAttendance` | Manage attendance of classes they control |
| `ManageResults` | Manage assessment results and submitted tasks for students of their groups and courses |
| `ManageCourses` | Modify the courses the teacher is in |
| `IsAdmin` | All of the above for any course or student |

Gamification, levels and experience are disabled for teachers.

## Notifications

Notifications are used for announcements and system messages related to module completion, submission review and gamification.

## Inventory and items

Students have an inventory that stores items. Teachers create items and assign them to course subjects. Some items can be exchanged for real-life prizes and merchandise.

- Some items are purchasable for money points; some can be converted to money points.
- Some kinds of items stack in a single slot; others do not.
- Inventory slots can be limited, and may be extended by other means, such as a special item.

## Achievements

Teachers create achievements and attach them to completion of course subjects. Achievements may reward money points and experience. Some achievements are generic and are awarded for statistics, such as total money or experience earned, daily streak and so on.
