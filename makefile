include app.env

cert:
	rm -rf certificates
	mkdir certificates
	openssl req -x509 -newkey rsa:4096 -keyout $(CERTIFICATE_PATH)/$(CERTIFICATE_KEY)  -out $(CERTIFICATE_PATH)/$(CERTIFICATE_FILE) -days 365 -nodes -subj "/CN=localhost"



test:
	go test -v -cover -short ./...

build: clean
	go build -o $(SRC_DIR)/$(EXE) 

clean:
	rm -r $(SRC_DIR)/$(EXE)


