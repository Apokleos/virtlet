/*
Copyright 2017 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// WriteFiles writes the files specified as a map under
// `targetDir`. The keys of the map are subpaths and values are file contents.
// WriteFiles automatically creates any non-existing directories mentioned in subpaths.
func WriteFiles(targetDir string, content map[string][]byte) error {
	for filename, bs := range content {
		fullPath := filepath.Join(targetDir, filename)
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error making directory %q: %v", dir, err)
		}
		if err := ioutil.WriteFile(fullPath, []byte(bs), 0644); err != nil {
			return fmt.Errorf("error writing %q: %v", fullPath, err)
		}
	}
	return nil
}

// GenIsoImage generates an ISO 9660 filesystem image containing
// files from srcDir. It uses specified volumeID as the volume id.
func GenIsoImage(isoPath string, volumeID string, srcDir string) error {
	out, err := exec.Command("genisoimage", "-o", isoPath, "-V", volumeID, "-r", "-J", srcDir).CombinedOutput()
	if err != nil {
		outStr := ""
		if len(out) != 0 {
			outStr = ". Output:\n" + string(out)
		}
		return fmt.Errorf("error generating iso: %v%s", err, outStr)
	}

	return nil
}
