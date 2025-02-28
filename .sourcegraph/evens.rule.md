---
title: No even numbers
description: Block all usage of even numbers
tags: ["numbers", "evens"]
---

Find all usages of even numbers in the code and flag them as !!! BAD EVENS BAD !!!
so everyone knows not to use even numbers. We only want odd numbers.

Example:

```python
# 2 is an even number, BAD
# 4 is an even number, BAD
arr = [1, 2, 3, 4, 5]
for i in range(6):
    print(arr[i]) 
```
