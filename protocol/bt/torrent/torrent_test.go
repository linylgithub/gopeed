package torrent

import (
	"github.com/monkeyWie/gopeed/protocol/bt/metainfo"
	"github.com/monkeyWie/gopeed/protocol/bt/peer"
	"testing"
)

func TestTorrent_Download(t *testing.T) {
	torrent := buildTorrent()
	torrent.Download("e:/testbt/download")
}

func buildTorrent() *Torrent {
	// metaInfo, err := metainfo.ParseFromFile("../testdata/Game.of.Thrones.S08E05.720p.WEB.H264-MEMENTO.torrent")
	metaInfo, err := metainfo.ParseFromFile("../testdata/office.torrent")
	if err != nil {
		panic(err)
	}

	metaInfo.AnnounceList[0] = append(metaInfo.AnnounceList[0], []string{
		"udp://tracker.coppersurfer.tk:6969/announce",
		"udp://tracker.leechers-paradise.org:6969/announce",
		"udp://tracker.opentrackr.org:1337/announce",
		"udp://p4p.arenabg.com:1337/announce",
		"udp://9.rarbg.to:2710/announce",
		"udp://9.rarbg.me:2710/announce",
		"udp://tracker.internetwarriors.net:1337/announce",
		"udp://exodus.desync.com:6969/announce",
		"udp://tracker.tiny-vps.com:6969/announce",
		"udp://tracker.sbsub.com:2710/announce",
		"udp://retracker.lanta-net.ru:2710/announce",
		"udp://open.stealth.si:80/announce",
		"udp://open.demonii.si:1337/announce",
		"udp://tracker.torrent.eu.org:451/announce",
		"udp://tracker.moeking.me:6969/announce",
		"udp://tracker.cyberia.is:6969/announce",
		"udp://denis.stalker.upeer.me:6969/announce",
		"udp://ipv4.tracker.harry.lu:80/announce",
		"udp://zephir.monocul.us:6969/announce",
		"udp://xxxtor.com:2710/announce",
	}...)
	return NewTorrent(peer.GenPeerID(), metaInfo)
}