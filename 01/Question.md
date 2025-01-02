# Question

Certainly! Here's a coding exercise focused on the Factory Design Pattern:

**Exercise: Implement a Factory Design Pattern**

Design a system where a `ShapeFactory` class is responsible for creating different types of shapes, such as `Circle`, `Square`, and `Rectangle`. Each shape should implement a common interface with a method `draw()`. The `ShapeFactory` should have a method `getShape(shapeType)` that returns an instance of the requested shape.

**Requirements:**

1. Define a `Shape` interface with a `draw()` method.
2. Implement `Circle`, `Square`, and `Rectangle` classes that conform to the `Shape` interface.
3. Create a `ShapeFactory` class with a `getShape(shapeType)` method that returns an object of type `Shape` based on the provided `shapeType`.
4. Demonstrate the usage of the factory by requesting different shapes and calling their `draw()` methods.

**Example Usage:**

```java
public class FactoryPatternDemo {
    public static void main(String[] args) {
        ShapeFactory shapeFactory = new ShapeFactory();

        // Get an object of Circle and call its draw method.
        Shape shape1 = shapeFactory.getShape("CIRCLE");
        shape1.draw();

        // Get an object of Square and call its draw method.
        Shape shape2 = shapeFactory.getShape("SQUARE");
        shape2.draw();

        // Get an object of Rectangle and call its draw method.
        Shape shape3 = shapeFactory.getShape("RECTANGLE");
        shape3.draw();
    }
}
```

This exercise helps in understanding how the Factory Design Pattern promotes loose coupling and scalability by delegating the responsibility of object instantiation to a factory class.

For a detailed explanation and implementation, you can refer to the following resource:

- [Factory Design Pattern in detail | Interview Question](https://www.youtube.com/watch?v=D5d1f9Lcmv4)

This video provides an in-depth understanding of the Factory Design Pattern with practical examples.

By working through this exercise, you can enhance your grasp of the Factory Design Pattern and its application in software design. 