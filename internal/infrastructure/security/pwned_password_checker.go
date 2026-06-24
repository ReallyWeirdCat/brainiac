// Copyright (c) 2026 Nikolai Papin
//
// This file is part of Brainiac gamification and education platform
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package security

import (
	"bufio"
	"os"
	"strings"
	"sync"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
)

// PwnedPasswordChecker checks passwords against a local rockyou.txt file.
type PwnedPasswordChecker struct {
	compromised       map[string]struct{}
	once              sync.Once
	minPasswordLength int
	filePath          string
	loadErr           error
}

var _ ports.CompromisedPasswordChecker = (*PwnedPasswordChecker)(nil)

func NewPwnedPasswordChecker(filePath string) *PwnedPasswordChecker {
	return &PwnedPasswordChecker{filePath: filePath}
}

func (p *PwnedPasswordChecker) load() {
	file, err := os.Open(p.filePath)
	if err != nil {
		p.loadErr = err
		return
	}
	defer file.Close()

	p.compromised = make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	// Increase buffer if lines are long (most passwords are short)
	const maxCapacity = 512 * 1024 // 512KB
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		pass := strings.TrimSpace(scanner.Text())
		// Skip empty lines
		if pass != "" && len(pass) > p.minPasswordLength {
			p.compromised[pass] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		p.loadErr = err
	}
}

func (p *PwnedPasswordChecker) IsCompromised(password string) (bool, error) {
	p.once.Do(p.load)
	if p.loadErr != nil {
		return false, p.loadErr
	}
	_, found := p.compromised[password]
	return found, nil
}
