#!/usr/bin/env sh

# Remove current fixtures
rm /app/assets/fixtures/*-output.txt

# Generate fixtures
# Help output
cowsay -h > /app/assets/fixtures/cow-help-output.txt 2>&1

# List output
cowsay -l > /app/assets/fixtures/cow-list-output.txt 2>&1

# Default flags output
cat /app/assets/fixtures/input.txt | cowsay 1> /app/assets/fixtures/cow-default-output.txt

# Small cow output
cat /app/assets/fixtures/input.txt | cowsay -f small.cow 1> /app/assets/fixtures/small-cow-output.txt

# Small cow borg eyes output
cat /app/assets/fixtures/input.txt | cowsay -f small.cow -b 1> /app/assets/fixtures/small-cow-borg-output.txt

# Small cow dead eyes output
cat /app/assets/fixtures/input.txt | cowsay -f small.cow -d 1> /app/assets/fixtures/small-cow-dead-output.txt

# Small cow greedy eyes output
cat /app/assets/fixtures/input.txt | cowsay -f small.cow -g 1> /app/assets/fixtures/small-cow-greedy-output.txt

# Small cow paranoid eyes output
cat /app/assets/fixtures/input.txt | cowsay -f small.cow -p 1> /app/assets/fixtures/small-cow-paranoid-output.txt

# Small cow stoned eyes output
cat /app/assets/fixtures/input.txt | cowsay -f small.cow -s 1> /app/assets/fixtures/small-cow-stoned-output.txt

# Small cow tired eyes output
cat /app/assets/fixtures/input.txt | cowsay -f small.cow -t 1> /app/assets/fixtures/small-cow-tired-output.txt

# Small cow wired eyes output
cat /app/assets/fixtures/input.txt | cowsay -f small.cow -w 1> /app/assets/fixtures/small-cow-wired-output.txt

# Small cow young eyes output
cat /app/assets/fixtures/input.txt | cowsay -f small.cow -y 1> /app/assets/fixtures/small-cow-young-output.txt

# Small cow custom eyes output
cat /app/assets/fixtures/input.txt | cowsay -f small.cow -e AA 1> /app/assets/fixtures/small-cow-custom-eyes-output.txt

# Small cow custom tongue output
cat /app/assets/fixtures/input.txt | cowsay -f small.cow -T AA 1> /app/assets/fixtures/small-cow-custom-tongue-output.txt

# Small cow extra wide output
cat /app/assets/fixtures/input.txt | cowsay -f small.cow -W 80 1> /app/assets/fixtures/small-cow-extra-wide-output.txt

# Small cow strip tabs and newlines output
cat /app/assets/fixtures/input-2.txt | cowsay -f small.cow 1> /app/assets/fixtures/small-cow-strip-tabs-and-newlines-output.txt

# Small cow expand tabs and keep newlines output
cat /app/assets/fixtures/input-2.txt | cowsay -f small.cow -n 1> /app/assets/fixtures/small-cow-expand-tabs-and-keep-newlines-output.txt

# Small cow extra wide short input output
echo "Test" | cowsay -f small.cow -W 80 1> /app/assets/fixtures/small-cow-extra-wide-short-input-output.txt
