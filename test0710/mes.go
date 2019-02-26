syntax = "proto3"

message SearchRequest {
	string query = 1;
	int32 page_number = 2;
	int32 result_per_page = 3;
	enum Corpus {
		UNIVERSAL = 0;
		WEB = 1;
		IMAGES = 2;
		LOACAL = 3;
		NEWS = 4;
}
	Corpus corpus = 4;


int main() {
}
