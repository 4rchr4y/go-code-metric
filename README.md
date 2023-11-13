## Lack of cohesion in methods(LCOM)

$$\text{LCOM} = 
\begin{cases} 
|P| - |Q|, & \text{if } |P| > |Q| \\
0, & \text{otherwise} 
\end{cases}$$

Where:

$P$ - number of pairs of methods in the class that do not share class fields

$Q$ - number of method pairs that share at least one class field

## Lack of cohesion in methods(LCOM96b)

$$LCOM96b = \frac{1}{a} \sum_{j=1}^{a} \left( \frac{m - \mu(A_j)}{m} \right)$$

$a$ - are number of attributes 

$m$ - are number of methods 

$\mu(A_j)$ - the number of methods accessing attribute $Aj$

## Distance from the main sequence(DMS)

$$ DMS = |A + I - 1| $$

Where:

$A$ - abstractness

$I$ - instability

## Instability

$$Instability = \frac{C^e}{C^e + C^a}$$

Where:

$C^e$ - number of efferent (or outgoing) dependencies

$C^a$ - number of afferent (or incoming) dependencies

## Abstractness

$$Abstractness = \frac{\sum m^a}{\sum m^c}$$

Where:

${\sum m^a}$ - sum of abstract elements (interfaces or abstract classes)

${\sum m^c}$ - sum of concrete elements (nonabstract classes, structs)

