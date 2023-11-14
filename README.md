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

![LCOM](assets/LCOM.webp)

Where:

$P$ - number of pairs of methods in the class that do not share class fields

$Q$ - number of method pairs that share at least one class field

## Lack of cohesion in methods(LCOM96b)

![LCOM96b](assets/LCOM96b.webp)

$a$ - are number of attributes 

$m$ - are number of methods 

$\mu(A_j)$ - the number of methods accessing attribute $Aj$

## Distance from the main sequence(DMS)

![DMS](assets/DMS.webp)

Where:

$A$ - abstractness

$I$ - instability

## Instability

![Instability](assets/instability.webp)

Where:

$C^e$ - number of efferent (or outgoing) dependencies

$C^a$ - number of afferent (or incoming) dependencies

## Abstractness

![Abstractness](assets/abstractness.webp)

Where:

${\sum m^a}$ - sum of abstract elements (interfaces or abstract classes)

${\sum m^c}$ - sum of concrete elements (nonabstract classes, structs)

## Average method complexity (AMC)

![AMC](assets/AMC.webp)

Where:

$MC$ - sum of the method complexities

$n$ - total number of methods in a class

## Cohesion among methods in class (CAM)

![CAM](assets/CAM.jpg)

Where:

$M_{shared}$ - number of method pairs that share attributes

$M_{total}$ - total number of possible method pairs

## Reusability index

![Reusability index](assets/reusability_index.jpg)

Where:

$C$ - cohesion

$K$ - coupling

$T$ - testability

$D$ - documentation

$w_1, w_2, w_3, w_4$ — weighting coefficients (the weighting coefficients are determined by the user)