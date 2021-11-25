Нужно разработать HTTP сервис для быстрого поиска анаграмм в словаре. \
Два слова считаются анаграммами, если одно можно получить из другого перестановкой букв (без учета регистра). \
Сервис должен предоставлять эндпоинт для загрузки списка слов в формате json. Пример использования: \
`curl localhost:17001/load -d '["foobar", "aabb", "baba", "boofar", "test"]'`

И эндпоинт для поиска анаграмм по слову в загруженном словаре. Примеры использования: \
`curl 'localhost:8080/get?word=foobar'  => ["foobar","boofar"]` \
`curl 'localhost:8080/get?word=raboof'  => ["foobar","boofar"]` \
`curl 'localhost:8080/get?word=abba'    => ["aabb","baba"]` \
`curl 'localhost:8080/get?word=test'    => ["test"]` \
`curl 'localhost:8080/get?word=qwerty'  => null`

# Решение
`make run` - запускает сервис локально на 17001 порту.
