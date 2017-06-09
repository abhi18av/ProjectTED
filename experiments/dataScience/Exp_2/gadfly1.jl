using Gadfly
using RDatasets

iris = dataset("datasets", "iris")



plot(iris, x=:SepalLength, y=:SepalWidth, color=:Species,
         Geom.point)

p = plot(iris, x=:SepalLength, y=:SepalWidth, Geom.histogram)
         #img = SVG("iris_plot.svg", 6inch, 4inch)
         #draw(img, p)

mammals = dataset("MASS", "mammals")
plot(mammals, x=:Body, y=:Brain, label=:Mammal, Geom.point, Geom.label)

gasoline = dataset("Ecdat", "Gasoline")
plot(gasoline, x=:Year, y=:LGasPCar, color=:Country,
         Geom.point, Geom.line)
