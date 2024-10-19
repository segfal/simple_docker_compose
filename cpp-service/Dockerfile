# Use an official GCC image
FROM gcc:latest

# Set working directory
WORKDIR /app

# Copy the C++ code
COPY . .

# Compile the C++ code
RUN g++ -o factorial main.cpp

# Run the compiled program
CMD ["./factorial"]