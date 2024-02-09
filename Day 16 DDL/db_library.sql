CREATE TABLE Buku(
	ID_Buku INT,
	Judul VARCHAR(255) NOT NULL,
	Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(ID_Buku)
);

CREATE TABLE Users (
	hp varchar(13) NOT NULL,
	nama varchar(50) NOT NULL,
	alamat varchar(255) NOT NULL,
	PRIMARY KEY(hp)
);

CREATE TABLE Genre(
	id_genre INT NOT NULL,
	nama_genre VARCHAR(50) NOT NULL,
	created_at DATE NOT NULL DEFAULT current_timestamp,
	PRIMARY KEY(id_genre)
);

CREATE TABLE Genre_Buku(
	ID_Buku INT,
	ID_Genre INT,
	FOREIGN KEY (ID_Buku) REFERENCES Buku(ID_Buku),
	FOREIGN KEY (ID_Genre) REFERENCES Genre(ID_Genre),
	PRIMARY KEY(ID_Buku, ID_Genre)
);

CREATE TABLE Peminjaman(
	hp varchar(13) NOT NULL,
	id_buku int NOT NULL,
	tanggal_pinjam date NOT NULL,
	tanggal_kembali date NOT NULL,
	FOREIGN KEY (id_buku) REFERENCES Buku(ID_Buku),
	FOREIGN KEY (hp) REFERENCES Users(hp),
	PRIMARY KEY (hp, id_buku)
);

--insert buku
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('14045','How to destroy people from inside', CURRENT_TIMESTAMP); 
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('14022','How god thinking about us', CURRENT_TIMESTAMP); 
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('91152','How to mastery javascript within 100 days', CURRENT_TIMESTAMP); 
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('17084','Dark story of Indonesia', CURRENT_TIMESTAMP); 
INSERT INTO BUKU (id_buku, judul, created_at) VALUES ('72902','Quantum of Supremacy by Michio Kaku', CURRENT_TIMESTAMP); 

--insert users
INSERT INTO users (hp,nama,alamat) VALUES ('085156385787','Fathur Nur Ihsan', 'Jln. Pondok Rumput 1 Gg. Bogo RT 01/12 No.21A');
INSERT INTO users (hp,nama,alamat) VALUES ('081294659238','Eka Agustianingsih', 'Cilebut Residence 1 Blok C-14 No.5');
INSERT INTO users (hp,nama,alamat) VALUES ('08998819826','John Tyson', 'BCC Blok G7 No.19');
INSERT INTO users (hp,nama,alamat) VALUES ('08998819405','Kevin Stevanus', 'Emerald Village Blok H-9 No.2');
INSERT INTO users (hp,nama,alamat) VALUES ('085693200789','Freddy Pasumain', 'Jln. Pegangsaan Timur No.56');

--insert genre
INSERT INTO genre (id_genre, nama_genre, created_at) VALUES ('001','Mystery', CURRENT_TIMESTAMP);
INSERT INTO genre (id_genre, nama_genre, created_at) VALUES ('002','Fantasy', CURRENT_TIMESTAMP);
INSERT INTO genre (id_genre, nama_genre, created_at) VALUES ('003','Action', CURRENT_TIMESTAMP);
INSERT INTO genre (id_genre, nama_genre, created_at) VALUES ('004','Shoujo', CURRENT_TIMESTAMP);
INSERT INTO genre (id_genre, nama_genre, created_at) VALUES ('005','Shounen', CURRENT_TIMESTAMP);

--insert genre_buku
INSERT INTO genre_buku (ID_Buku, ID_Genre) VALUES ('14045','003');
INSERT INTO genre_buku (ID_Buku, ID_Genre) VALUES ('14022','002');
INSERT INTO genre_buku (ID_Buku, ID_Genre) VALUES ('91152','005');
INSERT INTO genre_buku (ID_Buku, ID_Genre) VALUES ('17084','001');
INSERT INTO genre_buku (ID_Buku, ID_Genre) VALUES ('72902','001');

--insert peminjaman
INSERT INTO peminjaman (hp, id_buku, tanggal_pinjam, tanggal_kembali) VALUES ('085156385787','14045','2023-05-07','2023-05-10');
INSERT INTO peminjaman (hp, id_buku, tanggal_pinjam, tanggal_kembali) VALUES ('081294659238','14022','2023-06-12','2023-06-15');
INSERT INTO peminjaman (hp, id_buku, tanggal_pinjam, tanggal_kembali) VALUES ('08998819826','91152','2023-07-15','2023-07-18');
INSERT INTO peminjaman (hp, id_buku, tanggal_pinjam, tanggal_kembali) VALUES ('08998819405','17084','2023-08-19','2023-08-24');
INSERT INTO peminjaman (hp, id_buku, tanggal_pinjam, tanggal_kembali) VALUES ('085693200789','72902','2023-09-27','2023-09-30');


