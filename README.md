# Package design quality metrics

### Table of contents:

* [ LCOM ](#lack-of-cohesion-in-methodslcom)
* [ LCOM96b ](#lack-of-cohesion-in-methodslcom96b)
* [ DMS ](#distance-from-the-main-sequencedms)
* [ Instability ](#instability)
* [ Abstractness ](#abstractness)
* [ AMC ](#average-method-complexity-amc)
* [ CAM ](#cohesion-among-methods-in-class-cam)
* [ Reusability index ](#reusability-index)

Cohesion among methods in class (CAM)
Reusability index

## Lack of cohesion in methods(LCOM)

![LCOM](assets/LCOM.jpg)

Where:

P - number of pairs of methods in the class that do not share class fields

Q - number of method pairs that share at least one class field

## Lack of cohesion in methods(LCOM96b)

![LCOM96b](assets/LCOM96b.jpg)

a - are number of attributes 

m - are number of methods 

μ(Aj) - the number of methods accessing attribute Aj

## Distance from the main sequence(DMS)

![DMS](assets/DMS.jpg)

Where:

A - abstractness

I - instability

## Instability

![Instability](assets/instability.jpg)

Where:

C<sup>e</sup> - number of efferent (or outgoing) dependencies

C<sup>a</sup>- number of afferent (or incoming) dependencies

## Abstractness

![Abstractness](assets/abstractness.jpg)

Where:

∑ m<sup>a</sup> - sum of abstract elements (interfaces or abstract classes)

∑ m<sup>c</sup> - sum of concrete elements (nonabstract classes, structs)

## Average method complexity (AMC)

![AMC](assets/AMC.jpg)

Where:

∑ <sub>MC</sub> - sum of the complexity of all methods in the class

n - total number of methods in a class

## Cohesion among methods in class (CAM)

![CAM](assets/CAM.jpg)

Where:

M <sub>shared</sub> - number of method pairs that share attributes

M <sub>total</sub>- total number of possible method pairs

## Reusability index

![Reusability index](assets/reusability_index.jpg)

Where:

C - cohesion

K - coupling

T - testability

D - documentation


w<sub>1</sub>, w<sub>2</sub>, w<sub>3</sub>, w<sub>4</sub> — weighting coefficients (the weighting coefficients are determined by the user)