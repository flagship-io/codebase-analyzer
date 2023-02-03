from flagship.app import Flagship
from flagship.config import Config

from flagship.app import Flagship
from flagship.config import Config
from flagship.handler import FlagshipEventHandler

class CustomEventHandler(FlagshipEventHandler):
    def __init__(self):
        FlagshipEventHandler.__init__(self)

    def on_log(self, level, message):
        print("Log >> " + message)
        pass

    def on_exception_raised(self, exception, traceback):
        FlagshipEventHandler.on_exception_raised(self, exception, traceback)
        print("Exception >> " + str(exception))
        pass


 Flagship.instance().start("your_env_id", "your_api_key"
        Config(event_handler=CustomEventHandler())

Flagship.instance().start("your_env_id", "your_api_key", Config())

visitor = Flagship.instance().create_visitor("user_#1234", {'isVip':True})
visitor.synchronize_modifications()

vip_feature_enabled = visitor.get_modification('showBtn', False)
vip_feature_color = visitor.get_modification('btnColor', "red")
vip_feature_size = visitor.get_modification('btnSize', 16)
