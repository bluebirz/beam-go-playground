provider "google" {
  project = "bluebirz-playground"
  region  = "eu-west1"
}

resource "google_pubsub_topic" "test-topic1" {
  name = "test-topic1"
}

resource "google_pubsub_subscription" "test-sub1" {
  name  = "test-sub1"
  topic = google_pubsub_topic.test-topic1.name
}

resource "google_pubsub_topic" "test-topic2" {
  name = "test-topic2"
}

resource "google_pubsub_subscription" "test-sub2" {
  name  = "test-sub2"
  topic = google_pubsub_topic.test-topic2.name
}
