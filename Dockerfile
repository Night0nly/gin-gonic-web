FROM golang:1.11.1 AS build-env
ENV APP_DIR ${GOPATH}/src/github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web
WORKDIR ${APP_DIR}
COPY ./ ${APP_DIR}
RUN ln -s ${APP_DIR}/bin/ /app
RUN make all

FROM alpine:latest
COPY --from=build-env /app/exec-file /usr/local/bin/exec-file
ENTRYPOINT ["/usr/local/bin/exec-file"]c



# Example
# Use an official Python runtime as a parent image
#FROM python:2.7-slim

# Set the working directory to /app
#WORKDIR /app

# Copy the current directory contents into the container at /app
#COPY . /app

# Install any needed packages specified in requirements.txt
#RUN pip install --trusted-host pypi.python.org -r requirements.txt

# Make port 80 available to the world outside this container
#EXPOSE 80

# Define environment variable
#ENV NAME World

# Run app.py when the container launches
#CMD ["python", "app.py"]
