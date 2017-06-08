#Pkg.add("PyPlot")

cd("/Users/eklavya/Projects/Polyglot/ProjectTED/experiments/dataScience/Exp_2/julia")

using PyPlot
x = linspace(0,2*pi,1000); y = sin(3*x + 4*cos(2*x))
plot(x, y, color="red", linewidth=2.0, linestyle="--")

using Gadfly
plot(x=rand(10), y=rand(10), Geom.line)

p1 = plot(x=[1,2,3], y=[4,5,6])
p2 = plot(x=[1,2,3], y=[6,7,8])
vstack(p1,p2)

p3 = plot(x=[5,7,8], y=[8,9,10])
p4 = plot(x=[5,7,8], y=[10,11,12])

# these two are equivalent
vstack(hstack(p1,p2),hstack(p3,p4))
gridstack([p1 p2; p3 p4])



#title("My great data", hstack(p3,p4))
