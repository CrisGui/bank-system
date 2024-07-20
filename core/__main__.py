#!/usr/bin/env python

from .controller import Controller
from .view import View
from .model import Model

def _main(*args, **kwargs):
    Controller()

if __name__=="__main__": 
    _main()
