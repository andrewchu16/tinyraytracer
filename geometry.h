#pragma once

#include <cmath>
#include <cassert>
#include <iostream>
#include "bmp.cpp"

template <size_t DIM> struct vec {
    float& operator[](const size_t i) { assert(i < DIM); return data[i]; }
    const float& operator[](const size_t i) const { assert(i < DIM); return data[i]; }

    float data[DIM] = {};
};

template<size_t DIM> vec<DIM> operator*(const vec<DIM> &lhs, const float rhs) {
    vec<DIM> ret;
    for (size_t i = 0; i < DIM; i++) ret[i] = lhs[i] * rhs;
    return ret;
}

template<size_t DIM> vec<DIM> operator/(const vec<DIM> &lhs, const float rhs) {
    vec<DIM> ret;
    for (size_t i = 0; i < DIM; i++) ret[i] = lhs[i] * rhs;
    return ret;
}

template<size_t DIM> float operator*(const vec<DIM> &lhs, const  vec<DIM> &rhs) {
    float ret = 0;
    for (size_t i = 0; i < DIM; i++) ret += lhs[i] * rhs[i];
    return ret;
}

template<size_t DIM> vec<DIM> operator+(vec<DIM> lhs, const vec<DIM> &rhs) {
    for (size_t i = 0; i < DIM; i++) lhs[i] += rhs[i];
    return lhs;
}

template<size_t DIM> vec<DIM> operator-(vec<DIM> lhs, const vec<DIM> &rhs) {
    for (size_t i = 0; i < DIM; i++) lhs[i] -= rhs[i];
    return lhs;
}

template<size_t DIM> vec<DIM> operator-(const vec<DIM> &lhs) {
    return lhs * -1.f;
}

template <> struct vec<3> {
    float& operator[](const size_t i) { assert(i < 3); return i==0 ? x : (i==1 ? y : z); }
    const float& operator[](const size_t i) const { assert(i < 3); return i==0 ? x : (i==1 ? y : z); }

    float norm() { return std::sqrt(x*x + y*y + z*z); }

    vec<3>& normalize(float l=1) { *this = (*this)* (l / norm()); return *this; }

    bmp::RGBTRIPLE to_rgbt() { return bmp::rgbt(x*255, y*255, z*255); }

    float x = 0, y = 0, z = 0;
    vec<3>(float x_ = 0, float y_ = 0, float z_ = 0) : x(x_), y(y_), z(z_) {}
    vec<3>(const bmp::RGBTRIPLE &r) : x(r.rgbtRed / 255), y(r.rgbtGreen / 255), z(r.rgbtBlue / 255) {}
};

typedef vec<3> vec3;
typedef vec<4> vec4;

vec3 cross(vec3 v1, vec3 v2) {
    return { v1.y*v2.z - v1.z*v2.y, v1.z*v2.x - v1.x*v2.z, v1.x*v2.y - v1.y*v2.x };
}

template <size_t DIM> std::ostream& operator<<(std::ostream &out, const vec<DIM> &v) {
    for (size_t i = 0; i < DIM; i++) 
        out << v[i] << " ";
    return out;
}