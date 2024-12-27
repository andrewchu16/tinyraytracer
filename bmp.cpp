#pragma once
#include <cstdint>
#include <vector>
#include <fstream>
#include <cassert>

#define PACK( __Declaration__) __Declaration__ __attribute__((__packed__))

namespace bmp {

typedef uint8_t BYTE;
typedef uint32_t DWORD;
typedef int32_t LONG;
typedef uint16_t WORD;

typedef PACK(struct {
    WORD   bfType;
    DWORD  bfSize;
    WORD   bfReserved1;
    WORD   bfReserved2;
    DWORD  bfOffBits;
}) BMPFILEHEADER;

typedef PACK(struct {
    DWORD  biSize;
    LONG   biWidth;
    LONG   biHeight;
    WORD   biPlanes;
    WORD   biBitCount;
    DWORD  biCompression;
    DWORD  biSizeImage;
    LONG   biXPelsPerMeter;
    LONG   biYPelsPerMeter;
    DWORD  biClrUsed;
    DWORD  biClrImportant;
}) BMPINFOHEADER;
 
typedef PACK(struct {
    BYTE rgbtBlue;
    BYTE rgbtGreen;
    BYTE rgbtRed;    
}) RGBTRIPLE;

RGBTRIPLE rgbt(BYTE r, BYTE g, BYTE b) { return (RGBTRIPLE) {b, g, r}; }

class bmp24 {
    private:
        BMPFILEHEADER file_header;
        BMPINFOHEADER info_header;
        std::vector<std::vector<RGBTRIPLE>> v;
        std::string file_name;

        bmp24() { } // i want you to initialize with my constructor ok
    public:
        // index the image by v[y][x]
        bmp24(const LONG &width, const LONG &height, std::string f_name="out.bmp") {
            file_header.bfType = 0x4d42;
            file_header.bfSize = sizeof(BMPFILEHEADER) + sizeof(BMPINFOHEADER);
            file_header.bfReserved1 = 0;
            file_header.bfReserved2 = 0;
            file_header.bfOffBits = sizeof(BMPFILEHEADER) + sizeof(BMPINFOHEADER);

            info_header.biSize = sizeof(BMPINFOHEADER);
            info_header.biWidth = width;
            info_header.biHeight = height;
            info_header.biPlanes = 0;
            info_header.biBitCount = 3 * 8;
            info_header.biCompression = 0;
            info_header.biSizeImage = 0;

            info_header.biXPelsPerMeter = 0;
            info_header.biYPelsPerMeter = 0;

            info_header.biClrUsed = 0;
            info_header.biClrImportant = 0;

            v.resize(height, std::vector<RGBTRIPLE>(width));

            file_name = f_name;
        }

        constexpr inline LONG width() { return info_header.biWidth; }
        constexpr inline LONG height() { return info_header.biHeight; }

        std::vector<RGBTRIPLE>& operator[](const LONG &x) { return v[x]; }

        void write_file() {
            std::ofstream fout = std::ofstream(file_name, std::ios::binary);
            assert(fout);

            int padding = (4 - (width() * sizeof(RGBTRIPLE)) % 4) % 4;

            fout.write((char *) &file_header, sizeof(BMPFILEHEADER));
            fout.write((char *) &info_header, sizeof(BMPINFOHEADER));

            for (int y = height() - 1; y >= 0; y--) {
                fout.write((char *) &v[y][0], sizeof(RGBTRIPLE) * width());

                for (int k = 0; k < padding; k++) fout.put(0x00);
            }

            fout.close();
        }
};
}
