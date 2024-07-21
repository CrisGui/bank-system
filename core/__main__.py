#!/usr/bin/env python
# from dotenv import load_dotenv

from .controller import CoreController
from .view import CoreView
from .model import CoreModel

def _main(*args, **kwargs):
    # load_dotenv()

    con = CoreController()
    view = CoreView()
    model = CoreModel()
    
    con.run()

if __name__=="__main__": 
    _main()
