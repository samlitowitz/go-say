#!/usr/bin/env sh

TEST_DIR=/tmp/test
OUTPUT_DIR="${TEST_DIR}/output"
EXPECTED_DIR="${TEST_DIR}/expected"

rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

# Compare output
# Help output
#cowsay -h > /app/assets/fixtures/docker.cow-help-output.txt 2>&1

# List output
go-say -l > $OUTPUT_DIR/cow-list-output.txt
diff $EXPECTED_DIR/cow-list-output.txt $OUTPUT_DIR/cow-list-output.txt

# Default flags output
cat $EXPECTED_DIR/input.txt | go-say > $OUTPUT_DIR/cow-default-output.txt
diff $EXPECTED_DIR/cow-default-output.txt $OUTPUT_DIR/cow-default-output.txt

# Small cow output
cat $EXPECTED_DIR/input.txt | go-say -f small.cow > $OUTPUT_DIR/small-cow-output.txt
diff $EXPECTED_DIR/small-cow-output.txt $OUTPUT_DIR/small-cow-output.txt

# Small cow borg eyes output
cat $EXPECTED_DIR/input.txt | go-say -f small.cow -b > $OUTPUT_DIR/small-cow-borg-output.txt
diff $EXPECTED_DIR/small-cow-borg-output.txt $OUTPUT_DIR/small-cow-borg-output.txt

# Small cow dead eyes output
cat $EXPECTED_DIR/input.txt | go-say -f small.cow -d > $OUTPUT_DIR/small-cow-dead-output.txt
diff $EXPECTED_DIR/small-cow-dead-output.txt $OUTPUT_DIR/small-cow-dead-output.txt

# Small cow greedy eyes output
cat $EXPECTED_DIR/input.txt | go-say -f small.cow -g > $OUTPUT_DIR/small-cow-greedy-output.txt
diff $EXPECTED_DIR/small-cow-greedy-output.txt $OUTPUT_DIR/small-cow-greedy-output.txt

# Small cow paranoid eyes output
cat $EXPECTED_DIR/input.txt | go-say -f small.cow -p > $OUTPUT_DIR/small-cow-paranoid-output.txt
diff $EXPECTED_DIR/small-cow-paranoid-output.txt $OUTPUT_DIR/small-cow-paranoid-output.txt

# Small cow stoned eyes output
cat $EXPECTED_DIR/input.txt | go-say -f small.cow -s > $OUTPUT_DIR/small-cow-stoned-output.txt
diff $EXPECTED_DIR/small-cow-stoned-output.txt $OUTPUT_DIR/small-cow-stoned-output.txt

# Small cow tired eyes output
cat $EXPECTED_DIR/input.txt | go-say -f small.cow -t > $OUTPUT_DIR/small-cow-tired-output.txt
diff $EXPECTED_DIR/small-cow-tired-output.txt $OUTPUT_DIR/small-cow-tired-output.txt

# Small cow wired eyes output
cat $EXPECTED_DIR/input.txt | go-say -f small.cow -w > $OUTPUT_DIR/small-cow-wired-output.txt
diff $EXPECTED_DIR/small-cow-wired-output.txt $OUTPUT_DIR/small-cow-wired-output.txt

# Small cow young eyes output
cat $EXPECTED_DIR/input.txt | go-say -f small.cow -y > $OUTPUT_DIR/small-cow-young-output.txt
diff $EXPECTED_DIR/small-cow-young-output.txt $OUTPUT_DIR/small-cow-young-output.txt

# Small cow custom eyes output
cat $EXPECTED_DIR/input.txt | go-say -f small.cow -e AA > $OUTPUT_DIR/small-cow-custom-eyes-output.txt
diff $EXPECTED_DIR/small-cow-custom-eyes-output.txt $OUTPUT_DIR/small-cow-custom-eyes-output.txt

# Small cow custom tongue output
cat $EXPECTED_DIR/input.txt | go-say -f small.cow -T AA > $OUTPUT_DIR/small-cow-custom-tongue-output.txt
diff $EXPECTED_DIR/small-cow-custom-tongue-output.txt $OUTPUT_DIR/small-cow-custom-tongue-output.txt

# Small cow extra wide output
cat $EXPECTED_DIR/input.txt | go-say -f small.cow -W 80 > $OUTPUT_DIR/small-cow-extra-wide-output.txt
diff $EXPECTED_DIR/small-cow-extra-wide-output.txt $OUTPUT_DIR/small-cow-extra-wide-output.txt

# Small cow strip tabs and newlines output
cat $EXPECTED_DIR/input-2.txt | go-say -f small.cow > $OUTPUT_DIR/small-cow-strip-tabs-and-newlines-output.txt
diff $EXPECTED_DIR/small-cow-strip-tabs-and-newlines-output.txt $OUTPUT_DIR/small-cow-strip-tabs-and-newlines-output.txt

# Small cow expand tabs and keep newlines output
cat $EXPECTED_DIR/input-2.txt | go-say -f small.cow -n > $OUTPUT_DIR/small-cow-expand-tabs-and-keep-newlines-output.txt
diff $EXPECTED_DIR/small-cow-expand-tabs-and-keep-newlines-output.txt $OUTPUT_DIR/small-cow-expand-tabs-and-keep-newlines-output.txt

# Small cow extra wide short input output
echo "Test" | go-say -f small.cow -W 80 > $OUTPUT_DIR/small-cow-extra-wide-short-input-output.txt
diff $EXPECTED_DIR/small-cow-extra-wide-short-input-output.txt $OUTPUT_DIR/small-cow-extra-wide-short-input-output.txt
