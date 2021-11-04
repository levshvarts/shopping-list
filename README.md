# shopping-list
Shopping List webapp with Prometheus metrics endpoint

This is an example app with prometheus metrics integrated.
This app runs on port 5656

# URLs
`http://0.0.0.0:5656/apple` - add one apple to cart
`http://0.0.0.0:5656/orange` - add one orange to cart
`http://0.0.0.0:5656/cereal` - add one box of cereal to cart

url parameter `count` in the url allows to add multiples to cart.
Example:
`http://0.0.0.0:5656/apple?count=5` - add 5 apples to cart

`http://0.0.0.0:5656/list` - show current content of the cart

`http://0.0.0.0:5656/metrics` - prometheus metrics output for the app

# Docker build
You can build the docker image with `docker build`, `docker tag`

# k8s deployment
image can be deployed with the deployment/service yml files.

# Prometheus
with prometheus operator running you can apply `ServiceMonitor.yml` to have prometheus monitor the app.
At this point metrics are available from prometheus:

`curl "http://0.0.0.0:9090/api/v1/query?query=apples_purchased"`