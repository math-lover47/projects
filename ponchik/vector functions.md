# Function Documentation

## Overview
This file provides utility functions for mathematical operations, primarily used in 2D and 3D vector computations. It defines helper functions like clamping, length calculation, vector normalization, and geometric intersections with shapes like spheres, boxes, and planes.

---

## Utility Functions

### `float clamp(float value, float min, float max)`
Clamps a value between a specified minimum and maximum range.

- **Parameters**:
  - `value`: The value to clamp.
  - `min`: The lower bound.
  - `max`: The upper bound.
- **Returns**: The clamped value.

---

### `double sign(double a)`
Determines the sign of a value.

- **Parameters**:
  - `a`: The input value.
- **Returns**: `1` if positive, `-1` if negative, `0` if zero.

---

### `double step(double edge, double x)`
Implements a step function.

- **Parameters**:
  - `edge`: The threshold.
  - `x`: The input value.
- **Returns**: `1.0` if `x > edge`, otherwise `0.0`.

---

## Vector Functions

### `float length(vec2 const& v)`
Calculates the length (magnitude) of a 2D vector.

- **Parameters**:
  - `v`: The input 2D vector.
- **Returns**: The scalar length of the vector.

---

### `float length(vec3 const& v)`
Calculates the length (magnitude) of a 3D vector.

- **Parameters**:
  - `v`: The input 3D vector.
- **Returns**: The scalar length of the vector.

---

### `vec3 norm(vec3 v)`
Normalizes a 3D vector (scales it to unit length).

- **Parameters**:
  - `v`: The input 3D vector.
- **Returns**: The normalized 3D vector.

---

### `float dot(vec3 const& a, vec3 const& b)`
Computes the dot product of two 3D vectors.

- **Parameters**:
  - `a`, `b`: Input 3D vectors.
- **Returns**: A scalar representing the dot product.

---

### `vec3 abs(vec3 const& v)`
Calculates the component-wise absolute value of a 3D vector.

- **Parameters**:
  - `v`: The input 3D vector.
- **Returns**: A 3D vector with absolute values of the components.

---

### `vec3 sign(vec3 const& v)`
Calculates the sign of each component of a 3D vector.

- **Parameters**:
  - `v`: The input 3D vector.
- **Returns**: A 3D vector with component signs.

---

### `vec3 step(vec3 const& edge, vec3 v)`
Implements a component-wise step function for 3D vectors.

- **Parameters**:
  - `edge`: Threshold values for each component.
  - `v`: Input vector.
- **Returns**: A 3D vector where each component is `1.0` if `v` exceeds `edge`, otherwise `0.0`.

---

### `vec3 reflect(vec3 rd, vec3 n)`
Reflects a vector around a normal.

- **Parameters**:
  - `rd`: Incident vector.
  - `n`: Normal vector.
- **Returns**: The reflected vector.

---

### Rotation Functions

#### `vec3 rotateX(vec3 a, double angle)`
Rotates a 3D vector around the X-axis.

- **Parameters**:
  - `a`: The input 3D vector.
  - `angle`: The angle of rotation (in radians).
- **Returns**: The rotated vector.

#### `vec3 rotateY(vec3 a, double angle)`
Rotates a 3D vector around the Y-axis.

- **Parameters**:
  - `a`: The input 3D vector.
  - `angle`: The angle of rotation (in radians).
- **Returns**: The rotated vector.

#### `vec3 rotateZ(vec3 a, double angle)`
Rotates a 3D vector around the Z-axis.

- **Parameters**:
  - `a`: The input 3D vector.
  - `angle`: The angle of rotation (in radians).
- **Returns**: The rotated vector.

---

## Shape Intersection Functions

### `vec2 sphere(vec3 ro, vec3 rd, float r)`
Computes intersections of a ray with a sphere.

- **Parameters**:
  - `ro`: Ray origin.
  - `rd`: Ray direction.
  - `r`: Radius of the sphere.
- **Returns**: A 2D vector containing the near and far intersection distances, or `-1.0` if no intersection.

---

### `vec2 box(vec3 ro, vec3 rd, vec3 boxSize, vec3& outNormal)`
Computes intersections of a ray with an axis-aligned bounding box (AABB).

- **Parameters**:
  - `ro`: Ray origin.
  - `rd`: Ray direction.
  - `boxSize`: Half-dimensions of the box.
  - `outNormal`: Output vector representing the box surface normal at the intersection point.
- **Returns**: A 2D vector containing the near and far intersection distances, or `-1.0` if no intersection.

---

### `float plane(vec3 ro, vec3 rd, vec3 p, float w)`
Computes the intersection of a ray with a plane.

- **Parameters**:
  - `ro`: Ray origin.
  - `rd`: Ray direction.
  - `p`: Plane normal.
  - `w`: Plane constant.
- **Returns**: The intersection distance, or a negative value if no intersection occurs.
