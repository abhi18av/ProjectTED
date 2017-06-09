using Gadfly
using RDatasets

iris = dataset("datasets", "iris")


p = plot(iris, x=:SepalLength, y=:SepalWidth, Geom.point)
