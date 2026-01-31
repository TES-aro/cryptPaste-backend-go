# alustus
Testauksen kohteena on cryptPasten [backend](https://github.com/TES-aro/cryptPaste-backend-go/tree/main)
Palvelimen on tarkoitus yksinkertaisessti tallentaa sille annettu teksti ja
noutaa avainarvoa vastaava teksti takaisin. Itse salaus tullaan hoitamaan
asiakkaan puolella.

Ohjelman kehitys tullaan tekemään BDDn innoittamana, mutta jokseenkin joustavasti.

# cryptPaste
Itse ohjelma voidaan jakaa kahteen pääosaan: ensimmäinen on HTTP palvelin, joka
vastaanottaa pyyntöjä clientiltä ja välittää nämä eteenpäin.
toinen osa on kommunikointi tietokannan kanssa.

## 01 HTTP palvelin
Ohjelma tulee kommunikoimaan verkkosivun kanssa käyttäen HTTP protokollaa.
viestit ovat puhdasta teksitä.

Testausta helpottaa se, että luku ja kirjoitus funktiot ovat erillään

### 01.1 tavoitteet
Vain POST ja GET ovat hyväksyttyjä toimintoja.
kaikki erikseen mainitsemattoman kontaktit palauttavat **vain** virhekoodin 405
POST on hyväksytty vain "juuri" osoitteeseen *foo.org/api*
GET on hyväksytty vain juuri + {id} osoitteeseen *foo.org/api/{id}*
IDn pituus on 32 merkkiä, kaikki muut palauttavat 404 virheen.

niin vastaus GET pyyntöön kuin POST pyynnön headerin täytyy sisältää
Content-Type: text/plain
POST pyynnön koko ei saa olla 655360 bittiä suurempi. (laiskasti laskettuna
1024 riviä 80 merkkiä per rivi ei edes nätti 2^n mutta meh.)



## 02 tietokanta

Ohjelma tulee käyttämään Posrgress tietokantaa, sillä minulla sattuu olemaan
yksi jo pystyssä. Paras vaihtoehto olisi SQL toteutukselle agnostinen palvelin,
mutta en omaa aikaisempaa kokemusta tietokantojen käytöstä Go:lla, joten
yksinkertaistetaan.

### 02.1 
Aloitetaan kartoittamalla tarkoituksenmukaiset toiminnat
