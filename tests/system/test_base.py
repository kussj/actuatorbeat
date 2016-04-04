from actuatorbeat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Actuatorbeat normally
        """
        self.render_config_template(
                path=os.path.abspath(self.working_dir) + "/log/*"
        )

        actuatorbeat_proc = self.start_beat()
        self.wait_until( lambda: self.log_contains("actuatorbeat is running"))
        exit_code = actuatorbeat_proc.kill_and_wait()
        assert exit_code == 0
