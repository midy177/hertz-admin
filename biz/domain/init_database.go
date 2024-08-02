/*
 * Copyright 2024 HertzAdmin Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 */

package domain

import "context"

type InitDatabase interface {
	InitDatabase(ctx context.Context) error
}
