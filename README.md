# worcestershire

Facebook data photos & videos date helper

`worcestershire` can be used to automate copying and file date scripting, for example:
```sh
worcestershire $takeout $takeout/photos_and_videos/your_videos.html | 
	while read -r vid date time; do
		name="FB_${date}_${time}.${file##*.}"
		cp -nv "$vid" "$name"
 		touch -m -t "$date $time" "$name"
 	 done
```
> [!NOTE]
> The example presented above should only be used as reference on how to utilize `worcestershire`.

### Usage
```
$ worcestershire <takeout> <takeout>photos_and_videos/your_videos.html
<takeout>/photos_and_videos/videos/XXXXXXXX_XXXXXXXXXXXXXXX_XXXXXXXXXX_n_XXXXXXXXXXXXXXX.mp4 YYYY-MM-DD HH:MM:SS
[...]
$ worcestershire fb fb/photos_and_videos/album/3.html
<takeout>/photos_and_videos/MobileUploads_XXXXXXXXXX/XXXXXXXX_XXXXXXXXXXXXXXX_XXXXXXXXXX_n_XXXXXXXXXXXXXXX.jpg YYYY-MM-DD HH:MM:SS
[...]
```
