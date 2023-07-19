from PIL import Image
import imagehash
import sys

im = Image.open(sys.argv[1])
if im.format == "PNG":
    # and is not RGBA
    if im.mode != "RGBA":
        im = im.convert("RGBA")

print(imagehash.whash(im))
