# change variable to lowercase
VARIABLE=${VARIABLE,,}

# to uppercase
VARIABLE=${VARIABLE^^}

# to capital
VARIABLE=${VARIABLE^}

# remove specific character from string (case insensitive)
VARIABLE=$(echo ${VARIABLE} | sed 's/<character>//i')

# grep with previous/next lines
cat test.log | grep "abc" -A1 # next 1 line
cat test.log | grep "abc" -B1 # previous 1 line

# remove trailing character
# remove trailing slash
VARIABLE=${VARIABLE%/}