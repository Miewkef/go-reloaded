создай папку 'texts', в нем текстовый файл
затем с помощью команды 'cd main' на папке main запусти команду 'go run . {имя читаемого файла}.txt {имя нужного файла}.txt'

---


Данная программа обрабатывает первую команду по мере появления на слове которая стоит до команды(знаки препинания игнорируються). Работает только с английскими буквами.
Программа не обработает команду если перед ней стоит команда.

Доступные команды:

(hex) - переводит число из шеснадцатеричной в десятичную; my num is 4f (hex) -> my num is 79

(bin) - переводит число из двоичной в десятичную; my num is 111 (bin) -> my num is 7

(cap) - "капитализирует" слово; my (cap) word -> My word

(low) - меняет все буквы в слове на нижний регистр; My WOrd (low) -> My word

(up) - меняет все буквы в слове на верхний регистр; My WOrd (up) -> My WOrd

Если после последних трех команд стоит запятая и цифра, то команда выполнится на стольких словах, сколько указано в команде.
Пример: watashi (cap) wa natsuki subaru (cap, 2) -> Watashi wa Natsuki Subaru

А также программа перемещает знаки препинания "по местам", и меняет артикль a если следующее слово наинается на: a, e, i, o, u, h
Пример: '  Between heaven and a earth , I'm an chosen one !  ! !  ' -> '  Between heaven and an earth, I'm an chosen one!!!'