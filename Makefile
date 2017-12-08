init:
	@docker-compose run --rm dep init
install:
	@docker-compose run --rm dep ensure
serve-dev:
	@dev_appserver.py ./appengine --host 0.0.0.0 --admin_host 0.0.0.0 --skip_sdk_update_check
run:
	@docker-compose up app
