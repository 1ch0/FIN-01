// cue vet album.json album.cue
import "time"

#Album: {
	artist: string
	title:  string
	date:   string
	date:   time.Format("2006-01-02")
}

album: #Album
