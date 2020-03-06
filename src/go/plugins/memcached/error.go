/*
** Zabbix
** Copyright (C) 2001-2020 Zabbix SIA
**
** This program is free software; you can redistribute it and/or modify
** it under the terms of the GNU General Public License as published by
** the Free Software Foundation; either version 2 of the License, or
** (at your option) any later version.
**
** This program is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
** GNU General Public License for more details.
**
** You should have received a copy of the GNU General Public License
** along with this program; if not, write to the Free Software
** Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
**/

package memcached

type zabbixError struct {
	err string
}

func (e zabbixError) Error() string {
	return e.err
}

var (
	errorInvalidParams     = zabbixError{"Invalid parameters."}
	errorCannotFetchData   = zabbixError{"Cannot fetch data."}
	errorCannotMarshalJSON = zabbixError{"Cannot marshal JSON."}
	errorUnsupportedMetric = zabbixError{"Unsupported metric."}
	errorEmptyResult       = zabbixError{"Empty result."}
	errorUnknownSession    = zabbixError{"Unknown session."}
)
