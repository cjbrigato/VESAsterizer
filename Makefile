RUN=go run cmd/vesasterizer/main.go

torus:
	$(RUN) -music examples/music/demo2.vtm -model examples/models/demo/torus.obj -charset unicode -mode solid -fps 60

cow:
	$(RUN) -music examples/music/demo2.vtm -model examples/models/cow.obj -charset unicode -mode solid -fps 60
