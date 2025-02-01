include app.env

cert:
	rm -rf certificates
	mkdir certificates
	#openssl genpkey -algorithm RSA -out certificates/private.key -aes256
	#openssl req -x509 -key $(CERTIFICATE_PATH)/$(CERTIFICATE_KEY) -out $(CERTIFICATE_PATH)/$(CERTIFICATE_FILE) -days 365
	#openssl x509 -in $(CERTIFICATE_PATH)/$(CERTIFICATE_FILE) -text -noout
	openssl req -x509 -newkey rsa:4096 -keyout $(CERTIFICATE_PATH)/$(CERTIFICATE_KEY)  -out $(CERTIFICATE_PATH)/$(CERTIFICATE_FILE) -days 365 -nodes -subj "/CN=localhost"

	#openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout $(CERTIFICATE_PATH)/$(CERTIFICATE_KEY) -out $(CERTIFICATE_PATH)/$(CERTIFICATE_FILE) \
    #-subj "/C=US/ST=State/L=City/O=Organization/OU=Department/CN=localhost"



test:
	go test -v -cover -short ./...

build: clean
	go build -o $(SRC_DIR)/$(EXE) 

clean:
	rm -r $(SRC_DIR)/$(EXE)


