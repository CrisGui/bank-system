#!/usr/bin/env python

from .utils import Singleton
from loguru import logger

class CoreController(metaclass=Singleton):
    def __init__(self):
        logger.info("Controller imported")

    def run(self):
        _ = self.load_envs()
        logger.debug("Running Controller")
    def load_envs(self):
        ...
