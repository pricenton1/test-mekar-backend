-- phpMyAdmin SQL Dump
-- version 4.5.1
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: 01 Mar 2021 pada 05.58
-- Versi Server: 10.1.19-MariaDB
-- PHP Version: 5.6.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_user`
--
CREATE DATABASE IF NOT EXISTS `db_user` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;
USE `db_user`;

-- --------------------------------------------------------

--
-- Struktur dari tabel `m_account`
--

CREATE TABLE `m_account` (
  `account_id` varchar(36) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `m_account`
--

INSERT INTO `m_account` (`account_id`, `email`, `password`) VALUES
('85d6d016-7cc8-4ccf-b6a8-d78714d8cef4', 'anton@gmail.com', '$2a$04$wwTjDKqA6.0OOzI2MQmDO./nx.VyqpA.1jWx6RRblUEHJZEdnvwr.');

-- --------------------------------------------------------

--
-- Struktur dari tabel `m_pekerjaan`
--

CREATE TABLE `m_pekerjaan` (
  `id_pekerjaan` varchar(36) NOT NULL,
  `label_pekerjaan` varchar(36) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `m_pekerjaan`
--

INSERT INTO `m_pekerjaan` (`id_pekerjaan`, `label_pekerjaan`) VALUES
('8219c7cc-30bb-11eb-b405-c85b766bafe8', 'PNS'),
('821e309f-30bb-11eb-b405-c85b766bafe8', 'Wirausaha'),
('82225b07-30bb-11eb-b405-c85b766bafe8', 'Wiraswasta');

-- --------------------------------------------------------

--
-- Struktur dari tabel `m_pendidikan`
--

CREATE TABLE `m_pendidikan` (
  `id_pendidikan` varchar(36) NOT NULL,
  `label_pendidikan` varchar(36) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `m_pendidikan`
--

INSERT INTO `m_pendidikan` (`id_pendidikan`, `label_pendidikan`) VALUES
('5ddbb91e-30bb-11eb-b405-c85b766bafe8', 'SD'),
('5de17471-30bb-11eb-b405-c85b766bafe8', 'SMP'),
('5de57c7c-30bb-11eb-b405-c85b766bafe8', 'SMA'),
('5dedbec5-30bb-11eb-b405-c85b766bafe8', 'Diploma'),
('5df1331d-30bb-11eb-b405-c85b766bafe8', 'Sarjana'),
('5df4cbfb-30bb-11eb-b405-c85b766bafe8', 'Magister'),
('5df8d140-30bb-11eb-b405-c85b766bafe8', 'Doktor');

-- --------------------------------------------------------

--
-- Struktur dari tabel `m_user`
--

CREATE TABLE `m_user` (
  `id_user` varchar(36) NOT NULL,
  `nik` varchar(16) NOT NULL,
  `username` varchar(45) NOT NULL,
  `tgl_lahir` date NOT NULL,
  `id_pekerjaan` varchar(36) NOT NULL,
  `id_pendidikan` varchar(36) NOT NULL,
  `user_status` int(11) DEFAULT '1',
  `created_date` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `m_user`
--

INSERT INTO `m_user` (`id_user`, `nik`, `username`, `tgl_lahir`, `id_pekerjaan`, `id_pendidikan`, `user_status`, `created_date`) VALUES
('8a95efea-e6d6-47c3-88e7-d57b03f9712e', '3209120503980005', 'UJI COBA', '1998-03-05', '82225b07-30bb-11eb-b405-c85b766bafe8', '5dedbec5-30bb-11eb-b405-c85b766bafe8', 1, '2020-12-01 09:41:42'),
('db179baa-559b-4967-b75c-9e316d6564af', '1234567891123', 'UJI COBA 2', '2000-03-09', '8219c7cc-30bb-11eb-b405-c85b766bafe8', '5df1331d-30bb-11eb-b405-c85b766bafe8', 1, '2021-02-28 21:55:05'),
('dd3542db-2666-405b-9b99-291b99e637b1', '12345678910', 'UJI COBA 3', '2000-03-08', '8219c7cc-30bb-11eb-b405-c85b766bafe8', '5df1331d-30bb-11eb-b405-c85b766bafe8', 0, '2021-02-28 21:49:58');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `m_account`
--
ALTER TABLE `m_account`
  ADD PRIMARY KEY (`account_id`),
  ADD UNIQUE KEY `email_UNIQUE` (`email`);

--
-- Indexes for table `m_pekerjaan`
--
ALTER TABLE `m_pekerjaan`
  ADD PRIMARY KEY (`id_pekerjaan`);

--
-- Indexes for table `m_pendidikan`
--
ALTER TABLE `m_pendidikan`
  ADD PRIMARY KEY (`id_pendidikan`);

--
-- Indexes for table `m_user`
--
ALTER TABLE `m_user`
  ADD PRIMARY KEY (`id_user`),
  ADD UNIQUE KEY `nik_UNIQUE` (`nik`),
  ADD KEY `fk_m_user_m_pekerjaan_idx` (`id_pekerjaan`),
  ADD KEY `fk_m_user_m_pendidikan1_idx` (`id_pendidikan`);

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `m_user`
--
ALTER TABLE `m_user`
  ADD CONSTRAINT `fk_m_user_m_pekerjaan` FOREIGN KEY (`id_pekerjaan`) REFERENCES `m_pekerjaan` (`id_pekerjaan`),
  ADD CONSTRAINT `fk_m_user_m_pendidikan1` FOREIGN KEY (`id_pendidikan`) REFERENCES `m_pendidikan` (`id_pendidikan`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
