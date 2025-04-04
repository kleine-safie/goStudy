### Build and run コマンド:

go build -o build && ./build

### Build with escape analysis verbose

go build -gcflags "-m -l" user

### Get wasp object

go tool objdump -s main.createUser user
