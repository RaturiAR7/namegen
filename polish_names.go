package namegen

var(
	polishMaleFirstNames = []string{
	"Adam", "Aleksander", "Andrzej", "Antoni", "Arkadiusz", "Artur", "Bartłomiej", "Bartosz", "Cezary", "Damian",
	"Dariusz", "Daniel", "Dawid", "Dominik", "Filip", "Grzegorz", "Hubert", "Jacek", "Jakub", "Jan",
	"Jarosław", "Jędrzej", "Jerzy", "Kacper", "Kamil", "Karol", "Krzysztof", "Łukasz", "Maciej", "Marcin",
	"Marek", "Mariusz", "Mateusz", "Michał", "Paweł", "Piotr", "Radosław", "Robert", "Sebastian", "Sławomir",
	"Stanisław", "Tomasz", "Waldemar", "Wiesław", "Wiktor", "Wojciech", "Zbigniew", "Zdzisław", "Adrian", "Albert",
	"Aleks", "Aleksy", "Alfred", "Amadeusz", "Ambroży", "Anatol", "Antek", "Anton", "Arkady", "Armand",
	"Arnold", "Artur", "August", "Benedykt", "Beniamin", "Bernard", "Błażej", "Bogdan", "Bogumił", "Bogusław",
	"Bolesław", "Borys", "Bruno", "Cezar", "Cyprian", "Czesław", "Damian", "Daniel", "Dawid", "Dionizy",
	"Dobromir", "Dominik", "Edward", "Emanuel", "Emil", "Erwin", "Eugeniusz", "Feliks", "Ferdynand", "Filip",
	"Franciszek", "Gabriel", "Gniewomir", "Gustaw", "Henryk", "Herbert", "Hubert", "Ignacy", "Igor", "Ireneusz",
	"Jacek", "Jacenty", "Jakub", "Jan", "Janusz", "Jeremi", "Jerzy", "Joachim", "Józef", "Julian",
	"Juliusz", "Kajetan", "Kamil", "Karol", "Kazimierz", "Klemens", "Kornel", "Krzysztof", "Leon", "Leszek",
	"Lucjan", "Łukasz", "Maciej", "Marcel", "Marek", "Marian", "Mariusz", "Mateusz", "Maurycy", "Michał",
	"Mieczysław", "Mikołaj", "Milan", "Nikodem", "Norbert", "Oskar", "Patryk", "Paweł", "Piotr", "Radosław",
	"Rafał", "Remigiusz", "Robert", "Ryszard", "Sebastian", "Seweryn", "Sławomir", "Stanisław", "Stefan", "Sylwester",
	"Tadeusz", "Teodor", "Tobiasz", "Tomasz", "Wacław", "Waldemar", "Wawrzyniec", "Wiktor", "Witold", "Wojciech",
	"Zbigniew", "Zbyszek", "Zenon", "Ziemowit", "Zygmunt", "Adrian", "Albert", "Aleks", "Aleksy", "Alfred",
	"Amadeusz", "Ambroży", "Anatol", "Antek", "Anton", "Arkady", "Armand", "Arnold", "Artur", "August",
	"Benedykt", "Beniamin", "Bernard", "Błażej", "Bogdan", "Bogumił", "Bogusław", "Bolesław", "Borys", "Bruno",
	"Cezar", "Cyprian", "Czesław", "Damian", "Daniel", "Dawid", "Dionizy", "Dobromir", "Dominik", "Edward",
	"Emanuel", "Emil", "Erwin", "Eugeniusz", "Feliks", "Ferdynand", "Filip", "Franciszek", "Gabriel", "Gniewomir",
	"Gustaw", "Henryk", "Herbert", "Hubert", "Ignacy", "Igor", "Ireneusz", "Jacek", "Jacenty", "Jakub",
	"Jan", "Janusz", "Jeremi", "Jerzy", "Joachim", "Józef", "Julian", "Juliusz", "Kajetan", "Kamil",
	"Karol", "Kazimierz", "Klemens", "Kornel", "Krzysztof", "Leon", "Leszek", "Lucjan", "Łukasz", "Maciej",
	"Marcel", "Marek", "Marian", "Mariusz", "Mateusz", "Maurycy", "Michał", "Mieczysław", "Mikołaj", "Milan",
	"Nikodem", "Norbert", "Oskar", "Patryk", "Paweł", "Piotr", "Radosław", "Rafał", "Remigiusz", "Robert",
	"Ryszard", "Sebastian", "Seweryn", "Sławomir", "Stanisław", "Stefan", "Sylwester", "Tadeusz", "Teodor", "Tobiasz",
	"Tomasz",
	}
	polishFemaleFirstNames = []string{
	"Adela", "Agnieszka", "Alicja", "Alina", "Amelia", "Anastazja", "Anna", "Antonina", "Barbara", "Beata",
	"Bożena", "Cecylia", "Dagmara", "Danuta", "Dominika", "Dorota", "Edyta", "Eliza", "Elżbieta", "Emilia",
	"Ewa", "Gabriela", "Grażyna", "Halina", "Hanna", "Iga", "Irena", "Irmina", "Iwona", "Jadwiga",
	"Joanna", "Jolanta", "Julia", "Justyna", "Kamila", "Karolina", "Katarzyna", "Kinga", "Klaudia", "Kornelia",
	"Krystyna", "Laura", "Lidia", "Liliana", "Lucja", "Magdalena", "Małgorzata", "Maria", "Marta", "Martyna",
	"Mieczysława", "Milena", "Monika", "Natalia", "Nina", "Ola", "Olga", "Patrycja", "Paulina", "Renata",
	"Sabina", "Sandra", "Sylwia", "Teresa", "Urszula", "Wanda", "Weronika", "Wiesława", "Zofia", "Żaneta",
	"Adrianna", "Agata", "Aleksandra", "Alicja", "Aneta", "Anita", "Anna", "Barbara", "Beata", "Bożena",
	"Cecylia", "Dagmara", "Danuta", "Dominika", "Dorota", "Edyta", "Eliza", "Elżbieta", "Emilia", "Ewa",
	"Gabriela", "Grażyna", "Halina", "Hanna", "Iga", "Irena", "Irmina", "Iwona", "Jadwiga", "Joanna",
	"Jolanta", "Julia", "Justyna", "Kamila", "Karolina", "Katarzyna", "Kinga", "Klaudia", "Kornelia", "Krystyna",
	"Laura", "Lidia", "Liliana", "Lucja", "Magdalena", "Małgorzata", "Maria", "Marta", "Martyna", "Mieczysława",
	"Milena", "Monika", "Natalia", "Nina", "Ola", "Olga", "Patrycja", "Paulina", "Renata", "Sabina",
	"Sandra", "Sylwia", "Teresa", "Urszula", "Wanda", "Weronika", "Wiesława", "Zofia", "Żaneta", "Ada",
	"Adela", "Agata", "Agnieszka", "Alicja", "Alina", "Amelia", "Anastazja", "Antonina", "Bernadeta", "Bianka",
	"Cecylia", "Celina", "Czesława", "Dagmara", "Danuta", "Dominika", "Dorota", "Edyta", "Eliza", "Elżbieta",
	"Emilia", "Ewa", "Felicja", "Gabriela", "Genowefa", "Grażyna", "Halina", "Hanna", "Iga", "Irena",
	"Irmina", "Iwona", "Jadwiga", "Joanna", "Jolanta", "Julia", "Justyna", "Kamila", "Karolina", "Katarzyna",
	"Kinga", "Klaudia", "Kornelia", "Krystyna", "Laura", "Lidia", "Liliana", "Lucja", "Magdalena", "Małgorzata",
	"Maria", "Marta", "Martyna", "Mieczysława", "Milena", "Monika", "Natalia", "Nina", "Ola", "Olga",
	"Patrycja", "Paulina", "Renata", "Sabina", "Sandra", "Sylwia", "Teresa", "Urszula", "Wanda", "Weronika",
	"Wiesława", "Zofia", "Żaneta", "Aldona", "Aleksandra", "Alicja", "Aneta", "Anita", "Anna", "Barbara",
	"Beata", "Bożena", "Cecylia", "Celina", "Czesława", "Dagmara", "Danuta", "Dominika", "Dorota", "Edyta",
	"Eliza", "Elżbieta", "Emilia", "Ewa", "Felicja", "Gabriela", "Genowefa", "Grażyna", "Halina", "Hanna",
	"Iga", "Irena", "Irmina", "Iwona", "Jadwiga", "Joanna", "Jolanta", "Julia", "Justyna", "Kamila",
	"Karolina", "Katarzyna", "Kinga", "Klaudia", "Kornelia", "Krystyna", "Laura", "Lidia", "Liliana", "Lucja",
	"Magdalena", "Małgorzata", "Maria", "Marta", "Martyna", "Mieczysława", "Milena", "Monika", "Natalia", "Nina",
	"Ola", "Olga", "Patrycja", "Paulina", "Renata", "Sabina", "Sandra", "Sylwia", "Teresa", "Urszula",
	"Wanda", "Weronika", "Wiesława", "Zofia", "Żaneta", "Ada", "Adela", "Agata", "Agnieszka", "Alicja",
	"Alina", "Amelia", "Anastazja", "Antonina", "Bernadeta", "Bianka", "Cecylia", "Celina", "Czesława", "Dagmara",
	"Danuta", "Dominika", "Dorota", "Edyta", "Eliza", "Elżbieta", "Emilia", "Ewa", "Felicja", "Gabriela",
	"Genowefa", "Grażyna", "Halina", "Hanna", "Iga", "Irena", "Irmina", "Iwona", "Jadwiga", "Joanna",
	"Jolanta", "Julia", "Justyna", "Kamila", "Karolina", "Katarzyna", "Kinga", "Klaudia", "Kornelia", "Krystyna",
	"Laura", "Lidia", "Liliana", "Lucja", "Magdalena", "Małgorzata", "Maria", "Marta", "Martyna", "Mieczysława",
	"Milena", "Monika", "Natalia", "Nina", "Ola", "Olga", "Patrycja", "Paulina", "Renata", "Sabina",
	"Sandra", "Sylwia", "Teresa", "Urszula", "Wanda", "Weronika", "Wiesława", "Zofia", "Żaneta",
}
polishLastNames = []string{
	"Adamczyk", "Adamski", "Bąk", "Baran", "Bartosz", "Bednarek", "Białas", "Bielecki", "Bielski", "Bogdan",
	"Brzeziński", "Chmielewski", "Czarnecki", "Czech", "Duda", "Dziedzic", "Gajos", "Głowacki", "Górka", "Grabowski",
	"Grzelak", "Jabłoński", "Jankowski", "Janowski", "Kaczmarek", "Kalinowski", "Kamiński", "Kowalczyk", "Kowalewski", "Kowalik",
	"Kozłowski", "Krawczyk", "Król", "Kucharski", "Kwiatkowski", "Laskowski", "Lewandowski", "Lipiński", "Maciejewski", "Majewski",
	"Malinowski", "Mazurek", "Mazurkiewicz", "Michalski", "Nowak", "Nowakowski", "Nowicki", "Pawłowski", "Piotrowski", "Rutkowski",
	"Sikora", "Sobczak", "Sokołowski", "Szymański", "Walczak", "Wojciechowski", "Wojcik", "Woźniak", "Zając", "Zalewski",
	"Zawisza", "Zieliński", "Zieliński", "Zych", "Żuk", "Żukowski", "Adamska", "Adamczak", "Bąk", "Baran",
	"Bartosz", "Bednarek", "Białas", "Bielecka", "Bielska", "Bogdan", "Brzezińska", "Chmielewska", "Czarnecka", "Czech",
	"Duda", "Dziedzic", "Gajos", "Głowacka", "Górka", "Grabowska", "Grzelak", "Jabłońska", "Jankowska", "Janowska",
	"Kaczmarek", "Kalinowska", "Kamińska", "Kowalczyk", "Kowalewska", "Kowalik", "Kozłowska", "Krawczyk", "Król", "Kucharska",
	"Kwiatkowska", "Laskowska", "Lewandowska", "Lipińska", "Maciejewska", "Majewska", "Malinowska", "Mazurek", "Mazurkiewicz", "Michalska",
	"Nowak", "Nowakowska", "Nowicka", "Pawłowska", "Piotrowska", "Rutkowska", "Sikora", "Sobczak", "Sokołowska", "Szymańska",
	"Walczak", "Wojciechowska", "Wojcik", "Woźniak", "Zając", "Zalewska", "Zawisza", "Zielińska", "Zielińska", "Zych",
	"Żuk", "Żukowska", "Adamska", "Adamczak", "Bąk", "Baran", "Bartosz", "Bednarek", "Białas", "Bielecka",
	"Bielska", "Bogdan", "Brzezińska", "Chmielewska", "Czarnecka", "Czech", "Duda", "Dziedzic", "Gajos", "Głowacka",
	"Górka", "Grabowska", "Grzelak", "Jabłońska", "Jankowska", "Janowska", "Kaczmarek", "Kalinowska", "Kamińska", "Kowalczyk",
	"Kowalewska", "Kowalik", "Kozłowska", "Krawczyk", "Król", "Kucharska", "Kwiatkowska", "Laskowska", "Lewandowska", "Lipińska",
	"Maciejewska", "Majewska", "Malinowska", "Mazurek", "Mazurkiewicz", "Michalska", "Nowak", "Nowakowska", "Nowicka", "Pawłowska",
	"Piotrowska", "Rutkowska", "Sikora", "Sobczak", "Sokołowska", "Szymańska", "Walczak", "Wojciechowska", "Wojcik", "Woźniak", "Zając",
	"Zalewska", "Zawisza", "Zielińska", "Zielińska", "Zych", "Żuk", "Żukowska", "Adamski", "Adamczak", "Bąk", "Baran",
	"Bartosz", "Bednarek", "Białas", "Bielecki", "Bielski", "Bogdan", "Brzeziński", "Chmielewski", "Czarnecki", "Czech", "Duda",
	"Dziedzic", "Gajos", "Głowacki", "Górka", "Grabowski", "Grzelak", "Jabłoński", "Jankowski", "Janowski", "Kaczmarek", "Kalinowski",
	"Kamiński", "Kowalczyk", "Kowalewski", "Kowalik", "Kozłowski", "Krawczyk", "Król", "Kucharski", "Kwiatkowski", "Laskowski", "Lewandowski",
	"Lipiński", "Maciejewski", "Majewski", "Malinowski", "Mazurek", "Mazurkiewicz", "Michalski", "Nowak", "Nowakowski", "Nowicki", "Pawłowski",
	"Piotrowski", "Rutkowski", "Sikora", "Sobczak", "Sokołowski", "Szymański", "Walczak", "Wojciechowski",
	}
)