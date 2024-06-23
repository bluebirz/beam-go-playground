go run main.go --runner=dataflow \
  --project=bluebirz-playground \
  --region=europe-west1 \
  --staging_location=gs://bluebirz-beam-dataflow-temp \
  --topic_1="test-topic1" --subscription_1="test-sub1" \
  --topic_2="test-topic2" --subscription_2="test-sub2"

