all:
	make -C news_article_service
	make -C user_service
clean:
	make -C news_article_service clean
	make -C user_service clean
lint:
	golint user_service
	golint news_article_service
grpc:
	make -C user_service_grpc 
	