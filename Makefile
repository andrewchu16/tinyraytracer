all: main


main: tinyraytracer.cpp geometry.h bmp.cpp
	g++ -o main $<
# -lm ?
clean:
	rm main
