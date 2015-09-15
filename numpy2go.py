import numpy as np

import ctypes



import numpy.ctypeslib as npct

numpy2go = npct.load_library("numpy2go", ".")

array_1d_double = npct.ndpointer(dtype=np.double, ndim=1, flags='CONTIGUOUS')

numpy2go.Test.restype = None
numpy2go.Test.argtypes = [array_1d_double, ctypes.c_int]

data = np.array([0.0, 1.0, 2.0])

print "Python says", data

numpy2go.Test(data, len(data))

