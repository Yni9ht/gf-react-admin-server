package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GDQCwgOYkACIgycDMXlienpqUX6UFovqzg/LzSElYGxskc/4c2M83mXfQQOvK/lyt7/PuLqgY3hRxg5Qxd1CL9KzPa/uzZp/e7dy7im8jesYArydBRpedH4YfnaIwtFs1a6NG6/aVHLvaVM6OjHW6WLlzixyS+qKLMx3v359syt15fzHDYQiD/cx8xzVurFrTOiJnM9Ore7XFwz8cURoe4qpmgnzsMRUmWhtTrTKqOu76m/mrKFTTPT5lHDCQ5h7Z5bjaynLuw9k7GmsXa/m44Ox9czon62UwKjZ5y7JXpDc8PNm9U3K39t3btZc/7WUzcfN75OCgs/PdVGR1o6fW3yzYrLmW+VzzLbpd1gZdSx73lT9EUvhXm7xMerS9hjtdLfrp374kVzjl1WvnyWzHffihuhnQ4POJVeuc5S+bbR462TLttqxsSExZHOGxi4THZu5dG79sx66Xf7ZtWnP+0vt6k+EjSNztt1pe3WGZu1B98yRNTKhW2WfBW7sExFO9bD81DJmxP8q5jTvzwufX69e13n8/mpnx8feShTElmqk3/DRcNEv9nr1OrmpTl7Ji+/ry962nF66PFPVsw6CUEJ396zCbz9XlU898bD+xc+WLPvfnvc/1j/2hW/DSuPCC/gS2FdkeHPekTvcOnGsENbrC7GNrw0PCxSkJDWa1Wo42zyN3OrQdXW+3lzejqCv3wL+XdHzt7jcK5lWfDXU+XTmXcvXTAz52H3/6/3atgWeez//6KrsO3pvbrl3doJnCZ3vkp+beRkijlvvU5wjbxRkXZLTs3E6NmWlq+PH2ZVnjc59nAAv1bFo2WxVhNWfHw7J8s0SGYV41nVP9O6F+9WiJ59Me7eMgv7CTYey9o9JWx0wyYxB81r2PTy/daDQf/+F67YebpoRtVMOeMOu7IEvQLv8zY/84quWZ3L91686cEB2U8Lavh2xRfYvT15+NmLiVP8nFYF91py3tiU8eBlv0cb3/7FsTdtLPMvid9ktHtW07vpk3WgXI7Le/k+Gad7m/xV9z98z8DA8P9/gDc7xykbzUWHmRgYnokwMMByAwODPlpuYEfkBnAGqO7RTwDpRlYT4M3IJMKMyE3IJoNyEwz8bwSRePMWwijsToEAAYb/jnzMDFgcxsoGkmdiYGLoZGBg8GEG8QABAAD//8juwpzrAwAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
