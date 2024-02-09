CREATE TABLE Buku(
	ID_Buku INT,
	Judul VARCHAR(255) NOT NULL,
	Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(ID_Buku)
);

CREATE TABLE Genre(
	ID_Genre SERIAL,
	Nama_Genre VARCHAR(50) NOT NULL,
	created_at DATE NOT NULL DEFAULT current_timestamp,
	PRIMARY KEY(ID_Genre)
);

CREATE TABLE Genre_Buku(
	ID_Buku INT,
	ID_Genre INT,
	FOREIGN KEY (ID_Buku) REFERENCES Buku(ID_Buku),
	FOREIGN KEY (ID_Genre) REFERENCES Genre(ID_Genre),
	PRIMARY KEY(ID_Buku, ID_Genre)
);

CREATE TABLE Users (
	hp varchar(13) NOT NULL,
	nama varchar(50) NOT NULL,
	alamat varchar(255) NOT NULL,
	PRIMARY KEY(hp)
);

CREATE TABLE Peminjaman(
	ID_Buku int NOT NULL,
	hp varchar(13) NOT NULL,
	tanggal_pinjam date,
	tanggal_kembali date,
	FOREIGN KEY (id_Buku) REFERENCES Buku(ID_Buku),
	FOREIGN KEY (hp) REFERENCES Users(hp),
	PRIMARY KEY (ID_Buku, hp)
);

--insert buku
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('14045','How to destroy people from inside', CURRENT_TIMESTAMP); 
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('14022','How god thinking about us', CURRENT_TIMESTAMP); 
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('91152','How to mastery javascript within 100 days', CURRENT_TIMESTAMP); 
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('17084','Dark story of Indonesia', CURRENT_TIMESTAMP); 
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('72902','Quantum of Supremacy by Michio Kaku', CURRENT_TIMESTAMP);
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('44399','Daily Life Immortal King', CURRENT_TIMESTAMP); 
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('20124','Natsume Yuujinchou', CURRENT_TIMESTAMP);
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('20162','Barakamon', CURRENT_TIMESTAMP); 
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('20102','Renchan', CURRENT_TIMESTAMP); 

--insert genre
INSERT INTO genre (id_genre, nama_genre, created_at) VALUES ('001','Mystery', CURRENT_TIMESTAMP);
INSERT INTO genre (id_genre, nama_genre, created_at) VALUES ('002','Fantasy', CURRENT_TIMESTAMP);
INSERT INTO genre (id_genre, nama_genre, created_at) VALUES ('003','Action', CURRENT_TIMESTAMP);
INSERT INTO genre (id_genre, nama_genre, created_at) VALUES ('004','Shoujo', CURRENT_TIMESTAMP);
INSERT INTO genre (id_genre, nama_genre, created_at) VALUES ('005','Shounen', CURRENT_TIMESTAMP);

--insert genre_buku
INSERT INTO Genre_buku (ID_Buku, ID_Genre) VALUES ('14045','003');
INSERT INTO Genre_buku (ID_Buku, ID_Genre) VALUES ('14022','002');
INSERT INTO Genre_buku (ID_Buku, ID_Genre) VALUES ('44399','002');
INSERT INTO Genre_buku (ID_Buku, ID_Genre) VALUES ('20192','002');
INSERT INTO Genre_buku (ID_Buku, ID_Genre) VALUES ('91152','005');
INSERT INTO Genre_buku (ID_Buku, ID_Genre) VALUES ('17084','001');
INSERT INTO Genre_buku (ID_Buku, ID_Genre) VALUES ('72902','001');
INSERT INTO Genre_buku (ID_Buku, ID_Genre) VALUES ('20124','001');
INSERT INTO Genre_buku (ID_Buku, ID_Genre) VALUES ('20162','001');
INSERT INTO Genre_buku (ID_Buku, ID_Genre) VALUES ('20102','001');

--insert users
INSERT INTO Users (hp,nama,alamat) VALUES ('085156385787','Fathur Nur Ihsan', 'Jln. Pondok Rumput 1 Gg. Bogo RT 01/12 No.21A');
INSERT INTO Users (hp,nama,alamat) VALUES ('081294659238','Eka Agustianingsih', 'Cilebut Residence 1 Blok C-14 No.5');
INSERT INTO Users (hp,nama,alamat) VALUES ('08998819826','John Tyson', 'BCC Blok G7 No.19');
INSERT INTO Users (hp,nama,alamat) VALUES ('08998819405','Kevin Stevanus', 'Emerald Village Blok H-9 No.2');
INSERT INTO Users (hp,nama,alamat) VALUES ('085693200789','Freddy Pasumain', 'Jln. Pegangsaan Timur No.56');
INSERT INTO Users (hp,nama,alamat) VALUES ('08581111111','Bagus Gunadi', 'PIK 2');
INSERT INTO Users (hp,nama,alamat) VALUES ('08522222222','Sasaki Kojiro', 'Perumnas Cakra Bogor');
INSERT INTO Users (hp,nama,alamat) VALUES ('08513333333','Cai yun', 'Bogor Nirwana Residence');


--insert peminjaman
INSERT INTO peminjaman (id_buku, hp, tanggal_pinjam, tanggal_kembali) VALUES ('14045','085156385787','2023-05-07','2023-05-10');
INSERT INTO peminjaman (id_buku, hp, tanggal_pinjam, tanggal_kembali) VALUES (,'14022','081294659238','2023-06-12','2023-06-15');
INSERT INTO peminjaman (id_buku, hp, tanggal_pinjam, tanggal_kembali) VALUES (,'91152','08998819826','2023-07-15','2023-07-18');
INSERT INTO peminjaman (id_buku, hp, tanggal_pinjam, tanggal_kembali) VALUES (,'17084','08998819405','2023-08-19','2023-08-24');
INSERT INTO peminjaman (id_buku, hp, tanggal_pinjam, tanggal_kembali) VALUES (,'72902','085693200789','2023-09-27','2023-09-30');

--d. show search result by judul buku using join
SELECT Buku.ID_Buku, Buku.Judul, Genre_Buku.ID_Genre
FROM Buku
INNER JOIN Genre_Buku ON Buku.ID_Buku = Genre_Buku.ID_Buku
ORDER BY judul;

--d. show search result by judul buku using where
SELECT Judul
FROM Buku
WHERE judul = 'Dark story of Indonesia'

--e. show search result order by genre_id buku using join
SELECT B.Judul, G.ID_Genre
FROM Buku B
JOIN Genre_Buku G ON B.ID_Buku = G.ID_Buku
ORDER BY G.ID_Genre;

--f.show 
--1.
INSERT INTO Peminjaman (id_buku, hp, tanggal_pinjam) VALUES ('14045','085156385787','2023-01-09');
INSERT INTO Peminjaman (id_buku, hp, tanggal_pinjam) VALUES ('91152','08998819405','2023-01-09');
INSERT INTO Peminjaman (id_buku, hp, tanggal_pinjam) VALUES ('72902','081294659238','2023-01-09');
INSERT INTO Peminjaman (id_buku, hp, tanggal_pinjam) VALUES ('17084','085156385787','2023-01-09');
--2.
INSERT INTO Peminjaman (id_buku, hp, tanggal_pinjam, tanggal_kembali) VALUES ('14022','085156385787','2023-01-09','2023-01-12');
--3.
INSERT INTO Peminjaman (id_buku, hp, tanggal_pinjam, tanggal_kembali) VALUES ('20192','085156385787','2023-05-07','2023-06-01');
--4.
INSERT INTO Peminjaman (id_buku, hp, tanggal_pinjam, tanggal_kembali) VALUES ('17084','081294659238','2024-01-02','2024-01-05');
--5.
INSERT INTO Peminjaman (id_buku, hp, tanggal_pinjam, tanggal_kembali) VALUES ('14022','08998819826','2024-02-03','2024-02-08');

--g. change status of books already borrowed from user 1 and that book remarked as "dikembalikan"
UPDATE Peminjaman
SET tanggal_kembali = '2023-01-16'
WHERE ID_Buku = '14045' AND hp = '085156385787';
SELECT * FROM Peminjaman;

--h.show users data and lsit of books is still currently borrowed/not yet given back Users;
SELECT U.hp, B.Judul, P.tanggal_pinjam, P.tanggal_kembali
FROM Users U
JOIN Peminjaman P ON U.hp = P.hp
JOIN Buku B ON B.ID_Buku = P.ID_Buku
WHERE P.tanggal_kembali IS NULL
ORDER BY P.tanggal_kembali DESC;
s
--i show books data who already returned by users
SELECT U.hp, B.Judul, P.tanggal_pinjam, P.tanggal_kembali
FROM Users U
JOIN Peminjaman P ON U.hp = P.hp
JOIN Buku B ON B.ID_Buku = P.ID_Buku
WHERE P.tanggal_kembali IS NOT NULL
ORDER BY P.tanggal_kembali DESC;

--j show books data who never borrowed by users
SELECT B.ID_Buku, B.Judul, G.Nama_Genre
FROM Buku B
JOIN Peminjaman P ON B.ID_Buku = P.ID_Buku
JOIN Genre G ON G.Nama_Genre = B.Judul
WHERE P.tanggal_pinjam IS NULL;

--k show users and how many book already borrowed
SELECT U.hp, U.nama, COUNT(P.ID_Buku) AS total_books_borrowed
FROM Users U
JOIN Peminjaman P ON U.hp = P.hp
JOIN Buku B ON P.ID_Buku = B.ID_Buku
JOIN Genre G ON G.Nama_Genre = B.Judul
GROUP BY U.hp, U.nama;

--l show users data who never borrowed book
SELECT U.hp, U.nama
FROM Users U
LEFT JOIN Peminjaman P ON U.hp = P.hp
LEFT JOIN Buku B ON P.ID_Buku = B.ID_Buku
LEFT JOIN Genre G ON G.Nama_Genre = B.Judul
WHERE P.ID_Buku IS NULL
GROUP BY U.hp, U.nama;

--m show users data who already borrowed book
SELECT U.hp, U.nama
FROM Users U
JOIN Peminjaman P ON U.hp = P.hp
JOIN Buku B ON P.ID_Buku = B.ID_Buku
JOIN Genre G ON G.Nama_Genre = B.Judul
GROUP BY U.hp, U.nama;

--n show users who most borrowed book
SELECT U.hp, U.nama, COUNT(P.ID_Buku) AS total_books_borrowed
FROM Users U
JOIN Peminjaman P ON U.hp = P.hp
JOIN Buku B ON P.ID_Buku = B.ID_Buku
JOIN Genre G ON G.Nama_Genre = B.Judul
GROUP BY U.hp, U.nama
ORDER BY Jumlah_Buku_Dipinjam DESC;

--o show genres data by how many books in each genre
SELECT G.ID_Genre, G.Nama_Genre, COUNT(B.ID_Buku) AS total_books
FROM Genre G
LEFT JOIN Buku B ON G.Nama_Genre = B.Judul
GROUP BY G.ID_Genre, G.Nama_Genre
ORDER BY total_books DESC;

--p show genre who most borrowed by user
SELECT G.ID_Genre, G.Nama_Genre, COUNT(P.ID_Buku) AS total_borrowed
FROM Genre G
JOIN Buku B ON G.Nama_Genre = B.Judul
JOIN Peminjaman P ON B.ID_Buku = P.ID_Buku
GROUP BY G.ID_Genre, G.Nama_Genre
ORDER BY total_borrowed DESC;

--q show users data, include books who borrowed, genre on those book borowed in one query statement
SELECT U.hp, U.nama, B.Judul AS Judul_Buku, G.Nama_Genre
FROM Users U
JOIN Peminjaman P ON U.hp = P.hp
JOIN Buku B ON P.ID_Buku = B.ID_Buku
JOIN Genre G ON G.ID_Genre = B.ID_Genre;





