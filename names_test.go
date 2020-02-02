package expressions

import (
	"testing"
)

func TestFullName(t *testing.T) {
	shouldMatch := []string{
		"Λdrøn",
		`Արամ Խաչատրյան`,
		`Johann Strauß`,
		`Vaqif Səmədoğlu`,
		`Вагиф Сәмәдоғлу`,
		`Heydər Əliyev`,
		`Һејдәр Әлијев`,
		`René Magritte`,
		`সুকুমার রায়`,
		`ព្រះ​ពុទ្ឋឃោសាចារ‌្យ​ជួន​ណាត`,
		`Céline Dion`,
		`ᓱᓴᓐ ᐊᒡᓗᒃᑲᖅ`,
		`ᏍᏏᏉᏯ`,
		`章子怡`,
		`王菲`,
		`Antonín Dvořák`,
		`Søren Hauch-Fausbøll`,
		`Søren Kierkegård`,
		`عبدالحليم حافظ`,
		`أم كلثوم`,
		`ብርሃነ ዘርኣይ`,
		`ኃይሌ ገብረሥላሴ`,
		`Mika Häkkinen`,
		`Gérard Depardieu`,
		`Jean Réno`,
		`Camille Saint-Saëns`,
		`Mylène Demongeot`,
		`François Truffaut`,
		`⠇⠕⠥⠊⠎ ⠃⠗⠁⠊⠇⠇⠑`,
		`ედუარდ შევარდნაძე`,
		`Jürgen Klinsmann`,
		`Walter Schultheiß`,
		`Γιώργος Νταλάρας`,
		`András Sütő`,
		`माधुरी दीक्षित`,
		`محسن مخملباف`,
		`Sinéad O’Connor`,
		`יהורם גאון`,
		`Fabrizio De André`,
		`久保田    利伸`,
		`林原 めぐみ`,
		`宮崎 駿`,
		`森鷗外`,
		`テクス・テクサン`,
		`이설희`,
		`안성기`,
		`심은하`,
		`솅조ᇰ`,
		`Trevor Żahra`,
		`Tor Åge Bringsværd`,
		`Herbjørn Sørebø`,
		`نصرت فتح علی خان`,
		`Nicómedes Santa Cruz`,
		`Lech Wałęsa`,
		`Amália Rodrigues`,
		`Olga Tañón`,
		`Pūblius Cornēlius Scīpiō Africānus`,
		`Михаил Горбачёв`,
		`Борис Гребенщиков`,
		`שלום עליכם`,
		`Áillohaš`,
		`Nils Aslak Valkeapää`,
		`Ľudovít Štúr`,
		`Frane Milčinski - Ježek`,
		`Björn Borg`,
		`Ἀρχιμήδης`,
		`صدر الدين عيني`,
		`Садриддин Айнӣ`,
		`சிவாஜி கனேசன்`,
		`舒淇`,
		`李安`,
		`ธงไชย แม็คอินไตย์`,
		`Ніна Матвієнко`,
		`Солижон Шарипов`,
		`Brad Pitt`,
		`Alisher ibn G'iyosiddin Navoiy`,
		`Trịnh Công Sơn`,
		`Đorđe Balašević`,
		`Ђорђе Балашевић`,
	}
	shouldFail := []string{
		`; -- this is bad`,
	}
	for _, name := range shouldMatch {
		if !ExpFullName.MatchString(name) {
			t.Errorf("%q is not valid\n", name)
		}
	}
	for _, name := range shouldFail {
		if ExpFullName.MatchString(name) {
			t.Errorf("%q should not be valid\n", name)
		}
	}

}
