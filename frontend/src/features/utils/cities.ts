const cities = [
  "Абаза",
  "Абакан",
  "Абдулино",
  "Абинск",
  "Агидель",
  "Агрыз",
  "Адыгейск",
  "Азнакаево",
  "Азов",
  "Ак-Довурак",
  "Аксай",
  "Алагир",
  "Алапаевск",
  "Алатырь",
  "Алдан",
  "Алейск",
  "Александров",
  "Александровск",
  "Александровск-Сахалинский",
  "Алексеевка",
  "Алексин",
  "Алзамай",
  "Алупка",
  "Алушта",
  "Альметьевск",
  "Амурск",
  "Анадырь",
  "Анапа",
  "Ангарск",
  "Андреаполь",
  "Анжеро-Судженск",
  "Анива",
  "Апатиты",
  "Апрелевка",
  "Апшеронск",
  "Арамиль",
  "Аргун",
  "Ардатов",
  "Ардон",
  "Арзамас",
  "Аркадак",
  "Армавир",
  "Армянск",
  "Арсеньев",
  "Арск",
  "Артём",
  "Артёмовск",
  "Артёмовский",
  "Архангельск",
  "Асбест",
  "Асино",
  "Астрахань",
  "Аткарск",
  "Ахтубинск",
  "Ачинск",
  "Аша",
  "Бабаево",
  "Бабушкин",
  "Бавлы",
  "Багратионовск",
  "Байкальск",
  "Баймак",
  "Бакал",
  "Баксан",
  "Балабаново",
  "Балаково",
  "Балахна",
  "Балашиха",
  "Балашов",
  "Балей",
  "Балтийск",
  "Барабинск",
  "Барнаул",
  "Барыш",
  "Батайск",
  "Бахчисарай",
  "Бежецк",
  "Белая Калитва",
  "Белая Холуница",
  "Белгород",
  "Белебей",
  "Белёв",
  "Белинский",
  "Белово",
  "Белогорск",
  "Белогорск",
  "Белозерск",
  "Белокуриха",
  "Беломорск",
  "Белоозёрский",
  "Белорецк",
  "Белореченск",
  "Белоусово",
  "Белоярский",
  "Белый",
  "Бердск",
  "Березники",
  "Берёзовский",
  "Берёзовский",
  "Беслан",
  "Бийск",
  "Бикин",
  "Билибино",
  "Биробиджан",
  "Бирск",
  "Бирюсинск",
  "Бирюч",
  "Благовещенск",
  "Благовещенск",
  "Благодарный",
  "Бобров",
  "Богданович",
  "Богородицк",
  "Богородск",
  "Боготол",
  "Богучар",
  "Бодайбо",
  "Бокситогорск",
  "Болгар",
  "Бологое",
  "Болотное",
  "Болохово",
  "Болхов",
  "Большой Камень",
  "Бор",
  "Борзя",
  "Борисоглебск",
  "Боровичи",
  "Боровск",
  "Бородино",
  "Братск",
  "Бронницы",
  "Брянск",
  "Бугульма",
  "Бугуруслан",
  "Будённовск",
  "Бузулук",
  "Буинск",
  "Буй",
  "Буйнакск",
  "Бутурлиновка",
  "Валдай",
  "Валуйки",
  "Велиж",
  "Великие Луки",
  "Великий Новгород",
  "Великий Устюг",
  "Вельск",
  "Венёв",
  "Верещагино",
  "Верея",
  "Верхнеуральск",
  "Верхний Тагил",
  "Верхний Уфалей",
  "Верхняя Пышма",
  "Верхняя Салда",
  "Верхняя Тура",
  "Верхотурье",
  "Верхоянск",
  "Весьегонск",
  "Ветлуга",
  "Видное",
  "Вилюйск",
  "Вилючинск",
  "Вихоревка",
  "Вичуга",
  "Владивосток",
  "Владикавказ",
  "Владимир",
  "Волгоград",
  "Волгодонск",
  "Волгореченск",
  "Волжск",
  "Волжский",
  "Вологда",
  "Володарск",
  "Волоколамск",
  "Волосово",
  "Волхов",
  "Волчанск",
  "Вольск",
  "Воркута",
  "Воронеж",
  "Ворсма",
  "Воскресенск",
  "Воткинск",
  "Всеволожск",
  "Вуктыл",
  "Выборг",
  "Выкса",
  "Высоковск",
  "Высоцк",
  "Вытегра",
  "Вышний Волочёк",
  "Вяземский",
  "Вязники",
  "Вязьма",
  "Вятские Поляны",
  "Гаврилов Посад",
  "Гаврилов-Ям",
  "Гагарин",
  "Гаджиево",
  "Гай",
  "Галич",
  "Гатчина",
  "Гвардейск",
  "Гдов",
  "Геленджик",
  "Георгиевск",
  "Глазов",
  "Голицыно",
  "Горбатов",
  "Горно-Алтайск",
  "Горнозаводск",
  "Горняк",
  "Городец",
  "Городище",
  "Городовиковск",
  "Гороховец",
  "Горячий Ключ",
  "Грайворон",
  "Гремячинск",
  "Грозный",
  "Грязи",
  "Грязовец",
  "Губаха",
  "Губкин",
  "Губкинский",
  "Гудермес",
  "Гуково",
  "Гулькевичи",
  "Гурьевск",
  "Гурьевск",
  "Гусев",
  "Гусиноозёрск",
  "Гусь-Хрустальный",
  "Давлеканово",
  "Дагестанские Огни",
  "Далматово",
  "Дальнегорск",
  "Дальнереченск",
  "Данилов",
  "Данков",
  "Дегтярск",
  "Дедовск",
  "Демидов",
  "Дербент",
  "Десногорск",
  "Джанкой",
  "Дзержинск",
  "Дзержинский",
  "Дивногорск",
  "Дигора",
  "Димитровград",
  "Дмитриев",
  "Дмитров",
  "Дмитровск",
  "Дно",
  "Добрянка",
  "Долгопрудный",
  "Долинск",
  "Домодедово",
  "Донецк",
  "Донской",
  "Дорогобуж",
  "Дрезна",
  "Дубна",
  "Дубовка",
  "Дудинка",
  "Духовщина",
  "Дюртюли",
  "Дятьково",
  "Евпатория",
  "Егорьевск",
  "Ейск",
  "Екатеринбург",
  "Елабуга",
  "Елец",
  "Елизово",
  "Ельня",
  "Еманжелинск",
  "Емва",
  "Енисейск",
  "Ермолино",
  "Ершов",
  "Ессентуки",
  "Ефремов",
  "Железноводск",
  "Железногорск",
  "Железногорск",
  "Железногорск-Илимский",
  "Жердевка",
  "Жигулёвск",
  "Жиздра",
  "Жирновск",
  "Жуков",
  "Жуковка",
  "Жуковский",
  "Завитинск",
  "Заводоуковск",
  "Заволжск",
  "Заволжье",
  "Задонск",
  "Заинск",
  "Закаменск",
  "Заозёрный",
  "Заозёрск",
  "Западная Двина",
  "Заполярный",
  "Зарайск",
  "Заречный",
  "Заречный",
  "Заринск",
  "Звенигово",
  "Звенигород",
  "Зверево",
  "Зеленогорск",
  "Зеленоградск",
  "Зеленодольск",
  "Зеленокумск",
  "Зерноград",
  "Зея",
  "Зима",
  "Златоуст",
  "Злынка",
  "Змеиногорск",
  "Знаменск",
  "Зубцов",
  "Зуевка",
  "Ивангород",
  "Иваново",
  "Ивантеевка",
  "Ивдель",
  "Игарка",
  "Ижевск",
  "Избербаш",
  "Изобильный",
  "Иланский",
  "Инза",
  "Иннополис",
  "Инсар",
  "Инта",
  "Ипатово",
  "Ирбит",
  "Иркутск",
  "Исилькуль",
  "Искитим",
  "Истра",
  "Ишим",
  "Ишимбай",
  "Йошкар-Ола",
  "Кадников",
  "Казань",
  "Калач",
  "Калач-на-Дону",
  "Калачинск",
  "Калининград",
  "Калининск",
  "Калтан",
  "Калуга",
  "Калязин",
  "Камбарка",
  "Каменка",
  "Каменногорск",
  "Каменск-Уральский",
  "Каменск-Шахтинский",
  "Камень-на-Оби",
  "Камешково",
  "Камызяк",
  "Камышин",
  "Камышлов",
  "Канаш",
  "Кандалакша",
  "Канск",
  "Карабаново",
  "Карабаш",
  "Карабулак",
  "Карасук",
  "Карачаевск",
  "Карачев",
  "Каргат",
  "Каргополь",
  "Карпинск",
  "Карталы",
  "Касимов",
  "Касли",
  "Каспийск",
  "Катав-Ивановск",
  "Катайск",
  "Качканар",
  "Кашин",
  "Кашира",
  "Кедровый",
  "Кемерово",
  "Кемь",
  "Керчь",
  "Кизел",
  "Кизилюрт",
  "Кизляр",
  "Кимовск",
  "Кимры",
  "Кингисепп",
  "Кинель",
  "Кинешма",
  "Киреевск",
  "Киренск",
  "Киржач",
  "Кириллов",
  "Кириши",
  "Киров",
  "Киров",
  "Кировград",
  "Кирово-Чепецк",
  "Кировск",
  "Кировск",
  "Кирс",
  "Кирсанов",
  "Киселёвск",
  "Кисловодск",
  "Клин",
  "Клинцы",
  "Княгинино",
  "Ковдор",
  "Ковров",
  "Ковылкино",
  "Когалым",
  "Кодинск",
  "Козельск",
  "Козловка",
  "Козьмодемьянск",
  "Кола",
  "Кологрив",
  "Коломна",
  "Колпашево",
  "Кольчугино",
  "Коммунар",
  "Комсомольск",
  "Комсомольск-на-Амуре",
  "Конаково",
  "Кондопога",
  "Кондрово",
  "Константиновск",
  "Копейск",
  "Кораблино",
  "Кореновск",
  "Коркино",
  "Королёв",
  "Короча",
  "Корсаков",
  "Коряжма",
  "Костерёво",
  "Костомукша",
  "Кострома",
  "Котельники",
  "Котельниково",
  "Котельнич",
  "Котлас",
  "Котово",
  "Котовск",
  "Кохма",
  "Красавино",
  "Красноармейск",
  "Красноармейск",
  "Красновишерск",
  "Красногорск",
  "Краснодар",
  "Краснозаводск",
  "Краснознаменск",
  "Краснознаменск",
  "Краснокаменск",
  "Краснокамск",
  "Красноперекопск",
  "Краснослободск",
  "Краснослободск",
  "Краснотурьинск",
  "Красноуральск",
  "Красноуфимск",
  "Красноярск",
  "Красный Кут",
  "Красный Сулин",
  "Красный Холм",
  "Кремёнки",
  "Кропоткин",
  "Крымск",
  "Кстово",
  "Кубинка",
  "Кувандык",
  "Кувшиново",
  "Кудрово",
  "Кудымкар",
  "Кузнецк",
  "Куйбышев",
  "Кукмор",
  "Кулебаки",
  "Кумертау",
  "Кунгур",
  "Купино",
  "Курган",
  "Курганинск",
  "Курильск",
  "Курлово",
  "Куровское",
  "Курск",
  "Куртамыш",
  "Курчалой",
  "Курчатов",
  "Куса",
  "Кушва",
  "Кызыл",
  "Кыштым",
  "Кяхта",
  "Лабинск",
  "Лабытнанги",
  "Лагань",
  "Ладушкин",
  "Лаишево",
  "Лакинск",
  "Лангепас",
  "Лахденпохья",
  "Лебедянь",
  "Лениногорск",
  "Ленинск",
  "Ленинск-Кузнецкий",
  "Ленск",
  "Лермонтов",
  "Лесной",
  "Лесозаводск",
  "Лесосибирск",
  "Ливны",
  "Ликино-Дулёво",
  "Липецк",
  "Липки",
  "Лиски",
  "Лихославль",
  "Лобня",
  "Лодейное Поле",
  "Лосино-Петровский",
  "Луга",
  "Луза",
  "Лукоянов",
  "Луховицы",
  "Лысково",
  "Лысьва",
  "Лыткарино",
  "Льгов",
  "Любань",
  "Люберцы",
  "Любим",
  "Людиново",
  "Лянтор",
  "Магадан",
  "Магас",
  "Магнитогорск",
  "Майкоп",
  "Майский",
  "Макаров",
  "Макарьев",
  "Макушино",
  "Малая Вишера",
  "Малгобек",
  "Малмыж",
  "Малоархангельск",
  "Малоярославец",
  "Мамадыш",
  "Мамоново",
  "Мантурово",
  "Мариинск",
  "Мариинский Посад",
  "Маркс",
  "Махачкала",
  "Мглин",
  "Мегион",
  "Медвежьегорск",
  "Медногорск",
  "Медынь",
  "Межгорье",
  "Междуреченск",
  "Мезень",
  "Меленки",
  "Мелеуз",
  "Менделеевск",
  "Мензелинск",
  "Мещовск",
  "Миасс",
  "Микунь",
  "Миллерово",
  "Минеральные Воды",
  "Минусинск",
  "Миньяр",
  "Мирный",
  "Мирный",
  "Михайлов",
  "Михайловка",
  "Михайловск",
  "Михайловск",
  "Мичуринск",
  "Могоча",
  "Можайск",
  "Можга",
  "Моздок",
  "Мончегорск",
  "Морозовск",
  "Моршанск",
  "Мосальск",
  "Москва",
  "Муравленко",
  "Мураши",
  "Мурино",
  "Мурманск",
  "Муром",
  "Мценск",
  "Мыски",
  "Мытищи",
  "Мышкин",
  "Набережные Челны",
  "Навашино",
  "Наволоки",
  "Надым",
  "Назарово",
  "Назрань",
  "Называевск",
  "Нальчик",
  "Нариманов",
  "Наро-Фоминск",
  "Нарткала",
  "Нарьян-Мар",
  "Находка",
  "Невель",
  "Невельск",
  "Невинномысск",
  "Невьянск",
  "Нелидово",
  "Неман",
  "Нерехта",
  "Нерчинск",
  "Нерюнгри",
  "Нестеров",
  "Нефтегорск",
  "Нефтекамск",
  "Нефтекумск",
  "Нефтеюганск",
  "Нея",
  "Нижневартовск",
  "Нижнекамск",
  "Нижнеудинск",
  "Нижние Серги",
  "Нижний Ломов",
  "Нижний Новгород",
  "Нижний Тагил",
  "Нижняя Салда",
  "Нижняя Тура",
  "Николаевск",
  "Николаевск-на-Амуре",
  "Никольск",
  "Никольск",
  "Никольское",
  "Новая Ладога",
  "Новая Ляля",
  "Новоалександровск",
  "Новоалтайск",
  "Новоаннинский",
  "Нововоронеж",
  "Новодвинск",
  "Новозыбков",
  "Новокубанск",
  "Новокузнецк",
  "Новокуйбышевск",
  "Новомичуринск",
  "Новомосковск",
  "Новопавловск",
  "Новоржев",
  "Новороссийск",
  "Новосибирск",
  "Новосиль",
  "Новосокольники",
  "Новотроицк",
  "Новоузенск",
  "Новоульяновск",
  "Новоуральск",
  "Новохопёрск",
  "Новочебоксарск",
  "Новочеркасск",
  "Новошахтинск",
  "Новый Оскол",
  "Новый Уренгой",
  "Ногинск",
  "Нолинск",
  "Норильск",
  "Ноябрьск",
  "Нурлат",
  "Нытва",
  "Нюрба",
  "Нягань",
  "Нязепетровск",
  "Няндома",
  "Облучье",
  "Обнинск",
  "Обоянь",
  "Обь",
  "Одинцово",
  "Озёрск",
  "Озёрск",
  "Озёры",
  "Октябрьск",
  "Октябрьский",
  "Окуловка",
  "Олёкминск",
  "Оленегорск",
  "Олонец",
  "Омск",
  "Омутнинск",
  "Онега",
  "Опочка",
  "Орёл",
  "Оренбург",
  "Орехово-Зуево",
  "Орлов",
  "Орск",
  "Оса",
  "Осинники",
  "Осташков",
  "Остров",
  "Островной",
  "Острогожск",
  "Отрадное",
  "Отрадный",
  "Оха",
  "Оханск",
  "Очёр",
  "Павлово",
  "Павловск",
  "Павловский Посад",
  "Палласовка",
  "Партизанск",
  "Певек",
  "Пенза",
  "Первомайск",
  "Первоуральск",
  "Перевоз",
  "Пересвет",
  "Переславль-Залесский",
  "Пермь",
  "Пестово",
  "Петров Вал",
  "Петровск",
  "Петровск-Забайкальский",
  "Петрозаводск",
  "Петропавловск-Камчатский",
  "Петухово",
  "Петушки",
  "Печора",
  "Печоры",
  "Пикалёво",
  "Пионерский",
  "Питкяранта",
  "Плавск",
  "Пласт",
  "Плёс",
  "Поворино",
  "Подольск",
  "Подпорожье",
  "Покачи",
  "Покров",
  "Покровск",
  "Полевской",
  "Полесск",
  "Полысаево",
  "Полярные Зори",
  "Полярный",
  "Поронайск",
  "Порхов",
  "Похвистнево",
  "Почеп",
  "Починок",
  "Пошехонье",
  "Правдинск",
  "Приволжск",
  "Приморск",
  "Приморск",
  "Приморско-Ахтарск",
  "Приозерск",
  "Прокопьевск",
  "Пролетарск",
  "Протвино",
  "Прохладный",
  "Псков",
  "Пугачёв",
  "Пудож",
  "Пустошка",
  "Пучеж",
  "Пушкино",
  "Пущино",
  "Пыталово",
  "Пыть-Ях",
  "Пятигорск",
  "Радужный",
  "Радужный",
  "Райчихинск",
  "Раменское",
  "Рассказово",
  "Ревда",
  "Реж",
  "Реутов",
  "Ржев",
  "Родники",
  "Рославль",
  "Россошь",
  "Ростов",
  "Ростов-на-Дону",
  "Рошаль",
  "Ртищево",
  "Рубцовск",
  "Рудня",
  "Руза",
  "Рузаевка",
  "Рыбинск",
  "Рыбное",
  "Рыльск",
  "Ряжск",
  "Рязань",
  "Саки",
  "Салават",
  "Салаир",
  "Салехард",
  "Сальск",
  "Самара",
  "Санкт-Петербург",
  "Саранск",
  "Сарапул",
  "Саратов",
  "Саров",
  "Сасово",
  "Сатка",
  "Сафоново",
  "Саяногорск",
  "Саянск",
  "Светлогорск",
  "Светлоград",
  "Светлый",
  "Светогорск",
  "Свирск",
  "Свободный",
  "Себеж",
  "Севастополь",
  "Северо-Курильск",
  "Северобайкальск",
  "Северодвинск",
  "Североморск",
  "Североуральск",
  "Северск",
  "Севск",
  "Сегежа",
  "Сельцо",
  "Семёнов",
  "Семикаракорск",
  "Семилуки",
  "Сенгилей",
  "Серафимович",
  "Сергач",
  "Сергиев Посад",
  "Сердобск",
  "Серов",
  "Серпухов",
  "Сертолово",
  "Сибай",
  "Сим",
  "Симферополь",
  "Сковородино",
  "Скопин",
  "Славгород",
  "Славск",
  "Славянск-на-Кубани",
  "Сланцы",
  "Слободской",
  "Слюдянка",
  "Смоленск",
  "Снежинск",
  "Снежногорск",
  "Собинка",
  "Советск",
  "Советск",
  "Советск",
  "Советская Гавань",
  "Советский",
  "Сокол",
  "Солигалич",
  "Соликамск",
  "Солнечногорск",
  "Соль-Илецк",
  "Сольвычегодск",
  "Сольцы",
  "Сорочинск",
  "Сорск",
  "Сортавала",
  "Сосенский",
  "Сосновка",
  "Сосновоборск",
  "Сосновый Бор",
  "Сосногорск",
  "Сочи",
  "Спас-Деменск",
  "Спас-Клепики",
  "Спасск",
  "Спасск-Дальний",
  "Спасск-Рязанский",
  "Среднеколымск",
  "Среднеуральск",
  "Сретенск",
  "Ставрополь",
  "Старая Купавна",
  "Старая Русса",
  "Старица",
  "Стародуб",
  "Старый Крым",
  "Старый Оскол",
  "Стерлитамак",
  "Стрежевой",
  "Строитель",
  "Струнино",
  "Ступино",
  "Суворов",
  "Судак",
  "Суджа",
  "Судогда",
  "Суздаль",
  "Сунжа",
  "Суоярви",
  "Сураж",
  "Сургут",
  "Суровикино",
  "Сурск",
  "Сусуман",
  "Сухиничи",
  "Сухой Лог",
  "Сызрань",
  "Сыктывкар",
  "Сысерть",
  "Сычёвка",
  "Сясьстрой",
  "Тавда",
  "Таганрог",
  "Тайга",
  "Тайшет",
  "Талдом",
  "Талица",
  "Тамбов",
  "Тара",
  "Тарко-Сале",
  "Таруса",
  "Татарск",
  "Таштагол",
  "Тверь",
  "Теберда",
  "Тейково",
  "Темников",
  "Темрюк",
  "Терек",
  "Тетюши",
  "Тимашёвск",
  "Тихвин",
  "Тихорецк",
  "Тобольск",
  "Тогучин",
  "Тольятти",
  "Томари",
  "Томмот",
  "Томск",
  "Топки",
  "Торжок",
  "Торопец",
  "Тосно",
  "Тотьма",
  "Трёхгорный",
  "Троицк",
  "Трубчевск",
  "Туапсе",
  "Туймазы",
  "Тула",
  "Тулун",
  "Туран",
  "Туринск",
  "Тутаев",
  "Тында",
  "Тырныауз",
  "Тюкалинск",
  "Тюмень",
  "Уварово",
  "Углегорск",
  "Углич",
  "Удачный",
  "Удомля",
  "Ужур",
  "Узловая",
  "Улан-Удэ",
  "Ульяновск",
  "Унеча",
  "Урай",
  "Урень",
  "Уржум",
  "Урус-Мартан",
  "Урюпинск",
  "Усинск",
  "Усмань",
  "Усолье",
  "Усолье-Сибирское",
  "Уссурийск",
  "Усть-Джегута",
  "Усть-Илимск",
  "Усть-Катав",
  "Усть-Кут",
  "Усть-Лабинск",
  "Устюжна",
  "Уфа",
  "Ухта",
  "Учалы",
  "Уяр",
  "Фатеж",
  "Феодосия",
  "Фокино",
  "Фокино",
  "Фролово",
  "Фрязино",
  "Фурманов",
  "Хабаровск",
  "Хадыженск",
  "Ханты-Мансийск",
  "Харабали",
  "Харовск",
  "Хасавюрт",
  "Хвалынск",
  "Хилок",
  "Химки",
  "Холм",
  "Холмск",
  "Хотьково",
  "Цивильск",
  "Цимлянск",
  "Циолковский",
  "Чадан",
  "Чайковский",
  "Чапаевск",
  "Чаплыгин",
  "Чебаркуль",
  "Чебоксары",
  "Чегем",
  "Чекалин",
  "Челябинск",
  "Чердынь",
  "Черемхово",
  "Черепаново",
  "Череповец",
  "Черкесск",
  "Чёрмоз",
  "Черноголовка",
  "Черногорск",
  "Чернушка",
  "Черняховск",
  "Чехов",
  "Чистополь",
  "Чита",
  "Чкаловск",
  "Чудово",
  "Чулым",
  "Чусовой",
  "Чухлома",
  "Шагонар",
  "Шадринск",
  "Шали",
  "Шарыпово",
  "Шарья",
  "Шатура",
  "Шахты",
  "Шахунья",
  "Шацк",
  "Шебекино",
  "Шелехов",
  "Шенкурск",
  "Шилка",
  "Шимановск",
  "Шиханы",
  "Шлиссельбург",
  "Шумерля",
  "Шумиха",
  "Шуя",
  "Щёкино",
  "Щёлкино",
  "Щёлково",
  "Щигры",
  "Щучье",
  "Электрогорск",
  "Электросталь",
  "Электроугли",
  "Элиста",
  "Энгельс",
  "Эртиль",
  "Югорск",
  "Южа",
  "Южно-Сахалинск",
  "Южно-Сухокумск",
  "Южноуральск",
  "Юрга",
  "Юрьев-Польский",
  "Юрьевец",
  "Юрюзань",
  "Юхнов",
  "Ядрин",
  "Якутск",
  "Ялта",
  "Ялуторовск",
  "Янаул",
  "Яранск",
  "Яровое",
  "Ярославль",
  "Ярцево",
  "Ясногорск",
  "Ясный",
  "Яхрома",
]

export default cities