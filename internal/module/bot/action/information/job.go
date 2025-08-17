package information

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/information"
)

// логика экшена (2 поста в неделю. 1 - Обязательный, 2 - Мотивация. Если есть какие-то дополнительные, то они должны отправиться между)

// Do Сценарий "Информации"
func (a *Action) Do(ctx context.Context, ps dto.PatientScenario) error {
	// логика экшена
	return nil
}

// отправка обязательной темы
// !!!!!(условие) отправка доп темы
// отправка мотивации
// !!!!!(условие) отправка подготовки к новому этапу

// надо получить последний отправленный пост
// посмотреть его тему
// на основе темы надо понять, какая тема должна идти следующая
// получить следующий пост по теме пост для пользователя

// GetNextPost логика получения след поста
func (a *Action) getNextPost(ctx context.Context, ps dto.PatientScenario) error {
	var sentPostID int64
	defer func() {
		// в дефере помечаем какой пост мы отправили
		err := a.informationDal.MarkPostSent(ctx, ps.PatientID, sentPostID)
		if err != nil {
			return
		}
	}()

	lastSentPost, err := a.informationDal.GetLastSentPost(ctx, ps.PatientID)
	if err != nil {
		return err
	}

	switch lastSentPost.PostsThemeID {
	case information.RequiredTheme:
		// тема последнего поста - обязательный контент
		sentPostID = a.routeLastRequiredTheme(ctx, ps)
	case information.MotivationTheme:
		// тема последнего поста - мотивация
		sentPostID = a.routeLastMotivationTheme(ctx, ps)
	case information.PreparingToSecondTheme:
		// тема последнего поста - подготовка ко второму этапу
		sentPostID = a.routeLastPreparingToSecondTheme(ctx, ps)
	default:
		// тема последнего поста была дополнительной, ее назначил врач
		sentPostID = a.routeLastAnotherTheme(ctx, ps)
	}

	return nil
}

// routeLastRequiredTheme выбор поста если крайняя тема была - "обязательный контент"
func (a *Action) routeLastRequiredTheme(ctx context.Context, ps dto.PatientScenario) error {
	// запрос дополнительных тем
	// если такие посты ЕСТЬ, то отправляем этот пост
	// если таких постов НЕТ, то отправляем мотивацию
}

// routeLastMotivationTheme выбор поста если крайняя тема была - "мотивация"
func (a *Action) routeLastMotivationTheme(ctx context.Context, ps dto.PatientScenario) error {
	// если прошло 2 месяца с момента запуска
	// ТО отправляем "подготовка ко второму этапу"
	// ИНАЧЕ пропускаем
	// завершаем сценарий
	// отправляем "обязательный контент"
}

// routeLastPreparingToSecondTheme выбор поста если крайняя тема была - "подготовка ко второму этапу"
func (a *Action) routeLastPreparingToSecondTheme(ctx context.Context, ps dto.PatientScenario) error {
	// отправляем "обязательный контент"
}

// routeLastAnotherTheme выбор поста если крайняя тема была какой-то дополнительной
func (a *Action) routeLastAnotherTheme(ctx context.Context, ps dto.PatientScenario) error {
	// отправляем мотивацию
}
