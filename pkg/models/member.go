package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"strconv"

	"golang.org/x/exp/slices"
)

const MEMBER_DB_PATH = "C:\\Users\\Anas Youssef\\Documents\\projects\\new-ui\\db\\members.json"

type Member struct {
	Id                  int              `json:"id"`
	Intermediate        string           `json:"intermediate"`
	FirstName           string           `json:"firstName"`
	LastName            string           `json:"lastName"`
	BirthDate           string           `json:"birthDate"`
	Gender              string           `json:"gender"`
	Vip                 string           `json:"vip"`
	Country             string           `json:"country"`
	Cin                 string           `json:"cin"`
	DeliveryDate        string           `json:"deliveryDate"`
	DeliveryLocation    string           `json:"deliveryLocation"`
	CinValidUntil       string           `json:"cinValidUntil"`
	Nationality         string           `json:"nationality"`
	MoroccanNationality string           `json:"moroccanNationality"`
	FullFatherName      string           `json:"fullFatherName"`
	FullMotherName      string           `json:"fullMotherName"`
	Profession          Profession       `json:"profession"`
	Correspondance      []Correspondance `json:"correspondance"`
}

func MemberShowAll() ([]Member, error) {
	var members []Member

	content, err := ioutil.ReadFile(MEMBER_DB_PATH)
	if err != nil {
		// log.Fatal("ERROR WHEN OPENING FILE: ", err)
		return members, errors.New(fmt.Sprintf("ERROR TRYING TO OPEN THE FILE: %v", err))
	}

	if err = json.Unmarshal(content, &members); err != nil {
		return members, errors.New(fmt.Sprintf("ERROR MARSHALING THE FILE: %v", err))
		// log.Fatal("ERROR MARSHALING THE FILE: ", err)
	}

	return members, nil
}

func GetMemberByID(id int) (Member, int, error) {
	members, err := MemberShowAll()
	if err != nil {
		return Member{}, -1, err
	}

	idx := slices.IndexFunc(members, func(m Member) bool {
		return m.Id == id
	})

	if idx == -1 {
		return Member{}, -1, errors.New(fmt.Sprintf("Could not find member with the given id: %d", id))
	}

	return members[idx], idx, nil
}

func MemberIndex(w http.ResponseWriter, r *http.Request) {
	members, err := MemberShowAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(members)
}

func MemberShow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	memberIDParam := ctx.Value("memberID").(string)

	memberID, err := strconv.Atoi(memberIDParam)
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	member, _, err := GetMemberByID(memberID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(member)
}

func MemberStore(w http.ResponseWriter, r *http.Request) {
	var member Member

	err := json.NewDecoder(r.Body).Decode(&member)

	if err != nil {
		http.Error(w, "Couldn't decode the request's body. Please check if the json format is valid!", http.StatusBadRequest)
		return
	}

	members, err := MemberShowAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newMembers := append(members, member)
	slices.SortFunc(newMembers, func(a, b Member) bool {
		return a.Id < b.Id
	})

	membersJson, err := json.MarshalIndent(newMembers, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile(MEMBER_DB_PATH, membersJson, fs.ModeAppend)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(member)
}

func MemberUpdate(w http.ResponseWriter, r *http.Request) {
	var member Member

	err := json.NewDecoder(r.Body).Decode(&member)

	if err != nil {
		http.Error(w, "Couldn't decode the request's body. Please check if the json format is valid!", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	memberIDParam := ctx.Value("memberID").(string)

	memberID, err := strconv.Atoi(memberIDParam)
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	foundMember, idx, err := GetMemberByID(memberID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	foundMember = member

	members, err := MemberShowAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newDeletedMembers := append(members[:idx], members[idx+1:]...)
	newMembers := append(newDeletedMembers, foundMember)
	slices.SortFunc(newMembers, func(a, b Member) bool {
		return a.Id < b.Id
	})

	membersJson, err := json.MarshalIndent(newMembers, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile(MEMBER_DB_PATH, membersJson, fs.ModeAppend)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(member)
}

func MemberDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	memberIDParam := ctx.Value("memberID").(string)

	memberID, err := strconv.Atoi(memberIDParam)
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	_, idx, err := GetMemberByID(memberID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	members, err := MemberShowAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newDeletedMembers := append(members[:idx], members[idx+1:]...)
	slices.SortFunc(newDeletedMembers, func(a, b Member) bool {
		return a.Id < b.Id
	})

	membersJson, err := json.MarshalIndent(newDeletedMembers, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile(MEMBER_DB_PATH, membersJson, fs.ModeAppend)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	// resp := make(map[string]string)
	// resp["message"] = fmt.Sprintf("%d deleted successfully.", memberID)

	// jsonResp, err := json.Marshal(resp)
	// if err != nil {
	//      http.Error(w, "Couldn't parse response to json.", http.StatusInternalServerError)
	//      return
	// }

	// w.Write(jsonResp)
}
