# Script for converting html posts from Medium export

for file in *.html; do
  pandoc -f html -t markdown "$file" > output/"${file%.*}.md"
done
