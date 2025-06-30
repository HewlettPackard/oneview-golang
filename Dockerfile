FROM golang:1.20

# Install Python 3 and pip
RUN apt-get update && apt-get install -y python3 python3-pip && apt-get clean

# Set working directory
WORKDIR /go/src/github.com/HewlettPackard/oneview-golang

# Copy everything and build
COPY . .

# (Optional) Install Python requirements if needed
# RUN pip3 install -r requirements.txt

RUN go build -o oneview-golang .

# Entrypoint or default command
CMD ["./oneview-golang"]
