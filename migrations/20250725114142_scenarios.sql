-- +goose Up
-- +goose StatementBegin

insert into scenarios (name, description, is_active)
values ('init', 'ИНИТ, Сценарий для врача, когда врач исполняет команду /init', false),                      -- 1
       ('metrics_start', 'Метрики, ознакомление со сценарием метрики, запускается 1 раз', false),            -- 2
       ('metrics_retry', 'Метрики, повторяющийся сценарий, дефолт раз в неделю', false),                     -- 3
       ('therapy', 'Терапия, наступает сразу после обучения', false),                                        -- 4
       ('education', 'Обучение, цепочка роликов, запускается после вступления пользователя в чат', false),   -- 5
       ('recommendation',
        'Рекомендации, запускается 1 раз. Суть в том, чтобы понять применил ли пользователь рекомендации после приема',
        false),                                                                                              -- 6
       ('information_start', 'Информация, ознакомление со сценарием информация, запускается 1 раз', false),  -- 7
       ('information_retry', 'Информация, повторяющийся сценарий, дефолт раз в неделю', false),              -- 8
       ('lost', 'Потеряшка, запускается, когда пользователь долго не отвечает или раз в промежуток', false), -- 9
       ('control', 'Выведение на контроль, сценарий для врача', false); -- 10


insert into scenario_steps (scenario_id, step_order, content, is_final)
values (1, 1, 'Укажите Фамилию и Имя клиента', false),
       (1, 2, 'Укажите пол', false),
       (1, 3, 'Укажите дату рождения', false),
       (1, 4, 'Введите ссылку на метрики', true),
       (2, 1, 'Здравствуйте, {{.FirstName}}👋

Скажите, пожалуйста, вы уже заполнили первичные метрики на старте нашей работы?', false),
       (2, 2, 'Супер!🎉Вы молодец!

Следующие метрики нужно будет заполнить через неделю. Я напомню. 🤗  Посмотрим, какие будут результаты, как изменится ваше самочувствие, какие рекомендации вы начнете регулярно выполнять.',
        true),
       (2, 3, 'Очень важно заполнить метрики на старте терапии для оценки динамики вашего самочувствия в дальнейшем.

Ведь потом надо будет с чем-то сравнивать. Это займет всего 2-3 минуты. Лучше прямо сейчас, не откладывайте здоровье на потом. 😉

Уделите себе 2 минуты!', false),
       (2, 4, 'Вот ссылка на ваши метрики. Заполните, пожалуйста!😊
И я больше не буду приставать с напоминаниями😉

{{.MetricsLink}}', false),
       (3, 1, '{{.FirstName}}, здравствуйте! Подошло время заполнить метрики.', true),
       (3, 2, 'Вот ссылка на ваши метрики. Заполните, пожалуйста. Это займет, как обычно, всего 2-3 минуты😊

{{.MetricsLink}}', false),
       (3, 3, 'Получилось заполнить метрики?', false),
       (3, 4, 'Супер🎉Большое спасибо!

Я передаю врачу, он посмотрит.

Я вам напомню, когда подойдет время для заполнения метрик 🤖', false),
       (3, 5, 'Минутка занудства😉

Напомню, врачу (да и Вам тоже) очень важно оценивать динамику самочувствия, понимать, что у вас получается и не получается делать из рекомендаций. Это позволяет оценить результаты и скорректировать терпию, при необходимости.

Метрики - это СУПЕР ВАЖНЫЙ инструмент для отслеживания и закрепления результата. Заполните, пожалуйста.

Не понимаю, почему люди называют меня занудой, я же помогаю им🤷‍♂️', false),
       (3, 6, 'Когда будет свободная минутка - заполните, пожалуйста. Советую не откладывать, а то потом забудете. Почему бы не сделать это прямо сейчас?😉

{{.MetricsLink}}', true),
       (4, 1, 'Здравствуйте, {{.FirstName}}👋

Вы уже заказали доставку препаратов?
Не возникло трудностей?

Я готов помочь, если что🤗', false),
       (4, 2, 'Отлично! 🎉

Напишите врачу, когда заказ придет и вы начнете терапию.', false),
       (4, 3, 'Напишите, пожалуйста, ассистентам клиники. Они ответят на все вопросы и помогут вам сделать заказ.

Их контакт в телеге - @my_Optimum', false),
       (4, 4, 'Хорошо! Если будут вопросы, пожалуйста, напишите ассистентам клиники.

Их контакт в телеге - @my_Optimum', false),
       (4, 5, 'Хорошо! Если будут вопросы, пожалуйста, напишите ассистентам клиники.

Их контакт в телеге - @my_Optimum', false),
       (4, 6, '{{.FirstName}}, вы уже получили заказ? Начали терапию?', false),
       (4, 7, 'Ура, это отличная новость🎉

Принимайте терапию по рекомендациям вашего врача.

Если вдруг будут какие-то реакции со стороны желудка - сразу напишите врачу.', false),
       (4, 8, 'Отлично!🎉 Когда придут остальные препараты, вводите их в терапию.

Если вдруг будут какие-то реакции со стороны желудка - сразу напишите врачу.', false),
       (4, 9, 'Понятно. Тогда ожидаем. Надеюсь, ожидание будет недолгим😸', false),
       (4, 10, 'Здравствуйте, {{.FirstName}} !👋

Поделитесь, как проходит терапия? Не замечали нежелательные реакции?', false),
       (4, 11, 'Отлично!🎉Продолжайте терапию, все идет как надо!', false),
       (4, 12, 'Отлично!🎉Продолжайте терапию, все идет как надо!', false),
       (4, 13, '{{.FirstName}}, приветствую вас!👋

Пришли оставшиеся препараты? Уже начали принимать?', false),
       (4, 14, 'Отлично!😊Принимайте их, соблюдая рекомендации врача.', false),
       (4, 15, 'Ну что ж, ожидаем. Когда посылка придет - добавляйте новые препараты в терапию.', false),
       (4, 16, 'Здравствуйте, {{.FirstName}} !👋

Поделитесь, как проходит терапия? Не замечали нежелательные реакции?', false),
       (4, 17, 'Отлично🎉 Продолжайте терапию, все идет как надо!', false),
       (4, 18,
        '🔥Пожалуйста, сообщите вашему врачу об этих реакциях. Он расскажет вам, что можно сделать. Я ему сейчас тоже сообщу.',
        true),
       (5, 1,
        'Здравствуйте ! Я - трекер-бот, помощник врачей клиники myOptimum. Умею многое, и постараюсь быть вам полезен😉

Моя главная задача - сделать процесс восстановления здоровья удобным — чтобы ничего не выпало из поля зрения. Буду напоминать о важных событиях в терапии и рекомендациях.

Обратите внимание: я не врач,  и не могу вам давать медицинские советы. Если у вас есть вопросы к доктору, вы всегда можете ему написать личное сообщение.

📍Предлагаю пройти короткое, но очень важное обучение.

Отправлю вам 5 видео по 2-3 минуты. Посмотрите их, пожалуйста, чтобы все понять и получить максимальный эффект от программы!',
        false),
       (5, 2,
        'Посмотрите Видео 1.

Главный врач расскажет о сути предстоящей работы, как расставить акценты, чему уделить особое внимание.',
        false),
       (5, 3,
        'Видео 2.

Очень важное видео! Вы узнаете, как будет строиться работа и как правильно следовать терапии.',
        false),
       (5, 4,
        'Видео 3.

Как соблюдать все рекомендации? Как расставить приоритеты и выработать новые привычки?',
        false),
       (5, 5,
        'Видео 4.

Пожалуй, самое важное! Вы узнаете, как будет осуществляться сопровождение вашим врачом, что такое метрики и зачем их заполнять (спойлер: это МЕГАважно!)',
        false),
       (5, 6,
        'Давайте начнем!

Предлагаю вам прямо сейчас заполнить метрику. Переходите по ссылке и заполняйте в метрике первую колонку. Это займет всего 1-2 минуты. Советую не откладывать на потом, чтобы не забыть😉

{{.MetricsLink}}',
        false),
       (5, 7,
        'Заполнили метрики? Если еще не успели, сделайте это сейчас, нужно всего 1-2 минуты. Зато я больше не буду об этом спрашивать😉',
        false),
       (5, 8,
        'Отлично! Мне нравится, как мы с вами работаем🤗

Остался последний видеоролик.',
        false),
       (5, 9,
        'Хорошо, я напомню вам об этом через пару дней.

А пока посмотрите, пожалуйста, видео. Осталось совсем немного, всего 1 ролик✨',
        false),
       (5, 10,
        'Видео 5 (последнее).

Вы узнаете про 7 самых распространенных ошибок наших клиентов. Нам жаль, что они повторяются из раза в раз😥 Но вы узнаете о них заранее и наверняка не повторите!',
        false),
       (5, 11,
        'Кстати, а вы уже подписались на наш закрытый канал для клиентов? Там нет рекламы и всякой шелухи, только полезная информация от врачей и немного философии 😉

Если не подписаны, то сделайте это прямо сейчас, переходите по ссылке:

https://t.me/myoptimum',
        false),
       (5, 12,
        'Ура🎉Обучение завершено!

Впереди много интересного🤗',
        false),
       (5, 13,
        'Сейчас вы можете перейти к заказу препаратов. Чем раньше вы это сделаете - тем раньше начнете восстанавливать здоровье.

И обязательно прочитайте рекомендации!',
        false),
       (5, 14,
        'Если у вас остались вопросы, то вы можете написать их прямо сейчас:

Организационные вопросы и заказ препаратов - контакт ассистентов клиники в телеграмм @my_Optimum. Медицинские вопросы - напишите в личку вашему врачу {{.DoctorUsername}}.

Желаю вам плодотворной работы и быстрого восстановления здоровья!

Будем на связи!',
        true),
       (6, 1,
        'Здравствуйте, {{.FirstName}}!👋

Вы уже смогли ознакомиться с рекомендациями?',
        false),
       (6, 2,
        'Отлично!🎉🎉🎉

Начинайте следовать всем рекомендациям, и вы начнете замечать изменения в самочувствии!

Напоминаю, раздел "персональные рекомендации", в документе, который вам выслал врач, содержит все самые важные аспекты для вас.

А по ссылкам вы можете перейти в большой раздел "Ключевые рекомендации по образу жизни". Ознакомьтесь с ним, когда будет свободная минутка. Там много полезной информации по всем сферам образа жизни.',
        false),
       (6, 3,
        'Отлично, вы молодец!🎉

Как говориться, дорогу осилит идущий!

Напомню, что в высланном вам файле указаны персональные рекомендации для вас. Их нужно выполнять в первую очередь. А по ссылкам вы можете перейти в большой раздел "Ключевые рекомендации по образу жизни". Ознакомьтесь с ним, когда будет свободная минутка. Там много полезной информации по всем сферам образа жизни.',
        false),
       (6, 4,
        'Обязательно найдите время и почитайте❣️❣️❣️

Назначенная вам терапия и рекомендации работают в комплексе. Для максимального эффекта важно и то и другое.',
        false),
       (6, 5, 'Желаю отличного самочувствия и настроения! У вас все получится! 🤗', false),
       (6, 6, '{{.FirstName}}, скажите, пожалуйста, вы уже ознакомились со всеми рекомендациями? Было время?', false),
       (6, 7, '{{.FirstName}}, обязательно найдите время и почитайте! Это важно❗

Назначенная вам терапия и рекомендации работают в комплексе. Для максимального эффекта важно и то и другое.', false),
       (6, 8, 'Здравствуйте, {{.FirstName}}🤗 Дочитали рекомендации?', false),
       (6, 9, 'Здравствуйте, Рекомендации - неотъемлемая часть восстановления здоровья, которая работает в комплексе с назначенной терапией.

Пожалуйста, прочитайте их. Хорошо?😉

Напомню, что приоритет к выполнению - только персональные рекомендации. Это даст максимальный эффект. Менять всю свою жизнь и переходить сразу на "Супер-ЗОЖ" не нужно.😀

Будут вопросы - пишите вашему доктору😊

Желаю вам успеха! Больше вас дергать по прочтению рекомендаций не буду😉 Но мы будем с вами регулярно заполнять метрики и смотреть как у вас получается выполнять рекомендации.',
        true),
       (7, 1, 'Здравствуйте, {{.FirstName}} !

Буду вам периодически отправлять важную информацию по вашему здоровью и советы по образу жизни. А еще - различные лайфхаки😉

Очень много реально важной и практичной информации содержатся в ваших рекомендациях, но, к сожалению, не все их читают😥

Поэтому важные лайфхаки буду высылать вам лично

Напомню, что у нас есть закрытая группа в телеграмм, для клиентов, без рекламы.

Подключайтесь: https://t.me/myoptimum

Итак, первый пост.',
        true),
       (8, 1, 'Здравствуйте, {{.FirstName}} ! У меня для вас очередная полезная и интересная информация, почитайте!

👍Ставьте под постами лайки, если информация была полезной! 👍',
        false),
       (8, 2, 'Текст поста по теме Обязательный',
        false),
       (8, 3, '👩‍⚕️А информацию ниже вам настоятельно рекомендовал посмотреть врач! 👩‍⚕️ Высылаю вам по его просьбе.',
        false),
       (8, 4, 'Текст поста по теме Другие, по 1 посту по теме',
        false),
       (8, 5, 'Важно не только знать что делать, но и как делать, поддерживать мотивацию.

Лайфхаки ниже как раз об этом! 🤗',
        false),
       (8, 6, 'Текст поста по теме Мотивация', false),
       (8, 7, 'Отправляю вам информацию о том, как можно еще укрепить свое здоровье.', false),
       (8, 8, 'Текст поста по теме Подготовка к новому этапу', false),
       (9, 1, 'Здравствуйте, {{.FirstName}}! Давно не общались с вами.

Скажите, пожалуйста, на каком этапе сейчас находитесь?',
        false),
       (9, 2, 'Отлично! Не пропадайте, пожалуйста😉

Вы можете делиться своими успехами, задавать вопросы, заполнять метрики, задавать вопросы врачу. Наша поддержка и готовность во всем помочь - одна из составляющих вашего успеха.',
        false),
       (9, 3, 'Прекрасно!🎉

Если есть вопросы, пишите врачу. И еще просьба - заполните, пожалуйста, метрики. Нам важно отслеживать ваш прогресс.

{{.MetricsLink}}',
        false),
       (9, 4, 'Хорошо, делайте контрольные замеры. 👌

Только не затягивайте, пожалуйста, очень важно увидеть результаты сразу после программы, чтобы их закрепить.

Любые вопросы вы можете задать ассистентам клиники, их контакт в телеграм:

@my_Optimum',
        false),
       (9, 5, 'Введите, пожалуйста, дату', false),
       (9, 6, 'Хорошо, записал. Я напишу вам ближе к этой дате.',
        false),
       (9, 7, 'Здравствуйте, {{.FirstName}}!

Удалось ли вам вернуться в процесс восстановления здоровья?

Может быть, есть вопросы?

Мы с врачом на связи🤗',
        false),
       (9, 8, 'Прекрасно🎉

Если есть вопросы, пишите врачу. И еще просьба - заполните, пожалуйста, метрики. Нам важно отслеживать ваш прогресс.

{{.MetricsLink}}',
        false),
       (9, 9, 'Понятно. Могу получить от вас обратную связь по какой причине? Нам важно это понимать!',
        false),
       (9, 10, 'Ок, отлично! 🎉🎉🎉 Рад вашим успехам.

Появятся вопросы или решите заняться чем-то новеньким (генетика, питание, мозг прокачать, биохакинг, ...) пишите!

Мы на связи!',
        false),
       (9, 11, 'Понимаю 😞😞😞

Такое сейчас время. Но помните, занимаясь здоровьем, вы не отнимаете у себя время, а наоборот, получаете больше времени за счет энергии и продуктивности!

Будет посвободнее - пишите! Продолжим!',
        false),
       (9, 12, 'Понимаю 😞

Цены сейчас на медицину дорогие. Если что, можете с врачом обсудить бюджет, думаю можно оставить только ключевые вещи.',
        false),
       (9, 13, 'Понял. Такое бывает в начале терапии, однако все можно отрегулировать с врачом.

Я сообщу ему об этом и вы сможете обсудить.',
        false),
       (9, 14, 'Понял. Очень редко, но такое бывает.

Тут важно в таких случаях разобраться - для этого мы собираем консилиум с главным врачом и другими врачами и они смотрят вашу историю и дадут рекомендации.

Безусловно это будет бесплатно!

Я сообщу вашему врачу, он напишет вам.',
        false),
       (9, 15, 'Опишите, пожалуйста, причину', false),
       (9, 16, 'Спасибо за ответ! Мы его проанализируем!', true),
       (10, 1, 'Здравствуйте, {{.FirstName}}👋', false)
;

insert into step_buttons (scenario, step, button_text)
values (1, 2, 'М'),
       (1, 2, 'Ж'),
       (4, 1, 'Да'),
       (4, 1, 'Возникли трудности'),
       (4, 1, 'Позже сделаю'),
       (4, 1, 'Сейчас не могу, через неделю'),
       (4, 6, 'Да'),
       (4, 6, 'Частично'),
       (4, 6, 'Нет ещё'),
       (4, 6, 'Нет, не делал заказ'),
       (4, 10, 'Все хорошо'),
       (4, 10, 'Есть реакции'),
       (4, 13, 'Да'),
       (4, 13, 'Нет'),
       (4, 16, 'Все хорошо'),
       (4, 16, 'Есть реакции'),
       (2, 1, 'Да'),
       (2, 1, 'Нет'),
       (3, 3, 'Да'),
       (3, 3, 'Позже заполню'),
       (5, 7, 'Да'),
       (5, 7, 'Позже'),
       (6, 1, 'Да'),
       (6, 1, 'В процессе'),
       (6, 1, 'Нет'),
       (6, 6, 'Да'),
       (6, 6, 'Нет'),
       (6, 8, 'Да'),
       (6, 8, 'Нет'),
       (9, 1, 'Я в процессе!'),
       (9, 1, 'Я возвращаюсь в процесс'),
       (9, 1, 'Пока не сдал конрольные анализы'),
       (9, 1, 'Вернусь примерно к дате...'),
       (9, 1, 'Остановил работу'),
       (9, 7, 'Да'),
       (9, 7, 'Вернусь примерно к дате...'),
       (9, 7, 'Не планирую продолжать'),
       (9, 9, 'Достиг цели'),
       (9, 9, 'Нет времени'),
       (9, 9, 'Пока нет финансов на это'),
       (9, 9, 'Были побочные реакции'),
       (9, 9, 'Не получил результата'),
       (9, 9, 'Другое (нажмите и опишите текстом ниже, подалуйста)')
;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
