
echo "Running tests..."
go test -v -cover -coverprofile=./tmp/c.out | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''

echo ""
echo "Create coverage reports"
go tool cover -html=./tmp/c.out -o ./tmp/coverage.html 

echo "Finished"
